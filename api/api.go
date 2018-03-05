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

func main() {
	stats.Open()
	defer stats.Close()

	router := mux.NewRouter()
	router.HandleFunc("/players/{id:[0-9]+}", getPlayerEndPoint).Methods("GET")
	router.HandleFunc("/events", getEventsEndPoint).Methods("GET").Queries("uscf_id", "{uscf_id:[0-9]+}")
	router.HandleFunc("/tournaments/{id}", getTournamentEndPoint).Methods("GET")
	router.HandleFunc("/tournaments/{id}/sections", getSectionEndPoint).Methods("GET")
	router.HandleFunc("/sections/{id}", getSectionCrossTableEndPoint).Methods("GET")
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
