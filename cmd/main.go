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
	//services.InsertFirstMovie()
	r := mux.NewRouter()
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMoviesId).Methods("GET")
	r.HandleFunc("/movies", handlers.Create).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")
	//PUT /api/users/{id}: Updates the details of an existing user based on the ID.
	//DELETE /api/users/{id}: Deletes a user from the database based on the ID.
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
