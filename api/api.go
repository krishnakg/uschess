package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"uschess/statsdb"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var stats statsdb.StatsDB

type RoundResult struct {
	PlayerId    int    `json:"playerId"`
	PlayerName  string `json:"playerName"`
	PlayerColor uint8  `json:"playerColor"`
	Result      string `json:"result,omitempty"`
}

type SectionPairings struct {
	SectionID string `json:"sectionId,omitempty"`
	// Map of PlayerId -> {Round -> RoundResult}.
	PlayerResults map[int]map[int]RoundResult `json:"player,omitempty"`
}

func main() {
	stats.Open()
	defer stats.Close()

	router := mux.NewRouter()
	router.HandleFunc("/players/{id:[0-9]+}", getPlayerEndPoint).Methods("GET")
	router.HandleFunc("/events", getEventsEndPoint).Methods("GET").Queries("uscf_id", "{uscf_id:[0-9]+}")
	router.HandleFunc("/tournaments/{id}", getTournamentEndPoint).Methods("GET")
	router.HandleFunc("/tournaments/{id}/sections", getSectionEndPoint).Methods("GET")
	router.HandleFunc("/sections/{id}", getSectionCrossTableEndPoint).Methods("GET")
	router.HandleFunc("/games/{id}", getGamesInSectionEndPoint).Methods("GET")
	router.HandleFunc("/playersearch/{query}", getPlayerSearchEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func getPlayerEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uscfID, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	player, err := stats.GetPlayer(uscfID)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(player)
}

func getTournamentEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tournament, err := stats.GetTournament(params["id"])
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(tournament)
}

func getSectionEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sections, err := stats.GetSectionInfo(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(sections)
}

func getEventsEndPoint(w http.ResponseWriter, r *http.Request) {
	uscfID, err := strconv.Atoi(r.FormValue("uscf_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	performances, err := stats.GetEvents(uscfID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(performances)
}

func getPlayerSearchEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	players, err := stats.GetPlayerSearchResult(params["query"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(players)
}

func getSectionCrossTableEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	results, err := stats.GetSectionResults(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(results)
}

func getGamesInSectionEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	games, err := stats.GetGames(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	sectionPairings := convertGamesToSectionPairings(games)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(sectionPairings.PlayerResults)
}

func convertGamesToSectionPairings(games []statsdb.Game) (sectionPairings SectionPairings) {
	// map of player id to their results.
	sectionPairings.PlayerResults = make(map[int]map[int]RoundResult)
	for _, game := range games {
		// This should not be set every time in this loop.
		sectionPairings.SectionID = game.SectionID

		// We will create the map entries the first time we see the keys.
		if _, ok := sectionPairings.PlayerResults[game.Player1ID]; !ok {
			sectionPairings.PlayerResults[game.Player1ID] = make(map[int]RoundResult)
		}
		if _, ok := sectionPairings.PlayerResults[game.Player2ID]; !ok {
			sectionPairings.PlayerResults[game.Player2ID] = make(map[int]RoundResult)
		}
		roundResults := gameToRoundResults(game)
		sectionPairings.PlayerResults[game.Player1ID][game.Round] = roundResults[1]
		sectionPairings.PlayerResults[game.Player2ID][game.Round] = roundResults[0]
	}
	return sectionPairings
}

// Convert game information into two results for each player. This will let UI clients
// to easily showcase performance against opponents easily.
func gameToRoundResults(game statsdb.Game) (roundResults [2]RoundResult) {
	var player2Result string
	switch game.Result {
	case "W":
		player2Result = "L"
	case "L":
		player2Result = "W"
	case "D":
		player2Result = "D"
	}
	// roundResults[0] will contain information on Player2's result against Player1.
	roundResults[0] = RoundResult{game.Player1ID, game.Player1Name, game.Player1Color, player2Result}
	// roundResults[1] will contain information on Player1's result against Player2.
	roundResults[1] = RoundResult{game.Player2ID, game.Player2Name, game.Player2Color, game.Result}
	return roundResults
}
