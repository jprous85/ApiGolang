package main

import (
	"github.com/gorilla/handlers"
	_ "github.com/gorilla/handlers"
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

	headers := handlers.AllowedHeaders([]string{"X-Request-With", "Content-Type"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(headers, methods, origins)(myRouting)))
}

func main() {
	handleRequests()
}
