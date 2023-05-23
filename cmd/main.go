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
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMoviesId).Methods("GET")
	r.HandleFunc("/movies", handlers.Create).Methods("POST")

	//PUT /api/users/{id}: Updates the details of an existing user based on the ID.
	//DELETE /api/users/{id}: Deletes a user from the database based on the ID.
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
