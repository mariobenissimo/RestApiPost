package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
)

func main() {
	db.InizializeDatabase()
	// inizializeService()
	r := mux.NewRouter()
	r.HandleFunc("/getMovies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMoviesId).Methods("GET")
	r.HandleFunc("/movies", handlers.Create).Methods("POST")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
