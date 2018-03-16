package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"net/http"
	"strconv"
	"uschess/statsdb"

	"github.com/golang/glog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var stats statsdb.StatsDB

func main() {
	portPtr := flag.String("port", "8080", "Port at which api server will listen to.")

	flag.Parse()
	defer glog.Flush()

	// Initialize database connection.
	stats.Open()
	defer stats.Close()

	router := mux.NewRouter()
	router.HandleFunc("/players/{id:[0-9]+}", getPlayerEndPoint).Methods("GET")
	router.HandleFunc("/events", getEventsEndPoint).Methods("GET").Queries("uscf_id", "{uscf_id:[0-9]+}")
	router.HandleFunc("/tournaments/", getTournamentListEndPoint).Methods("GET")
	router.HandleFunc("/tournaments/{id}", getTournamentEndPoint).Methods("GET")
	router.HandleFunc("/tournaments/{id}/sections", getSectionEndPoint).Methods("GET")
	router.HandleFunc("/sections/{id}", getSectionCrossTableEndPoint).Methods("GET")
	router.HandleFunc("/games/{id}", getGamesInSectionEndPoint).Methods("GET")
	router.HandleFunc("/playersearch/{query}", getPlayerSearchEndPoint).Methods("GET")

	glog.Info("Server starting..")
	glog.Fatal(http.ListenAndServe(":"+*portPtr, router))
	glog.Info("Shutting down..")
}

func getPlayerEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uscfID, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
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

func getTournamentListEndPoint(w http.ResponseWriter, r *http.Request) {
	tournaments, err := stats.GetTournaments(20)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(tournaments)
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
	sectionID := params["id"]
	games, err := stats.GetGames(sectionID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	sectionPairings := NewSectionPairings(sectionID)
	sectionPairings.addGames(games)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(sectionPairings.PlayerResults)
}
