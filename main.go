package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouting := mux.NewRouter().StrictSlash(true)
	myRouting.HandleFunc("/", index).Methods("GET")
	myRouting.HandleFunc("/", create).Methods("POST")
	myRouting.HandleFunc("/{id}", show).Methods("GET")
	myRouting.HandleFunc("/{id}", update).Methods("PUT")
	myRouting.HandleFunc("/{id}", delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouting))
}

func main() {
	handleRequests()
}
