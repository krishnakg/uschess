package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/player/{id}", GetPlayer).Methods("GET")
	http.ListenAndServe(":8080", router)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {

}
