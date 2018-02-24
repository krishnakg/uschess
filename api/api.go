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
	router.HandleFunc("/players/{id:[0-9]+}", GetPlayerEndPoint).Methods("GET")
	router.HandleFunc("/events", GetEventsEndPoint).Methods("GET").Queries("uscf_id", "{uscf_id:[0-9]+}")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPlayerEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uscfId, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	player, err := stats.GetPlayer(uscfId)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(player)
}

func GetEventsEndPoint(w http.ResponseWriter, r *http.Request) {
	uscfId, err := strconv.Atoi(r.FormValue("uscf_id"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	performances, err := stats.GetEvents(uscfId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(performances)
}
