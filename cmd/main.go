package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	"github.com/mariobenissimo/RestApiPost/pkg/logger"
	"github.com/mariobenissimo/RestApiPost/pkg/middleware"
)

func main() {
	db.InizializeDatabase()
	logger.IniziazeLogger()
	r := mux.NewRouter()
	r.Use(middleware.HeaderMiddleware)
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.JwtVerify)
	r.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	s.HandleFunc("/movies", handlers.Create).Methods("POST")
	s.HandleFunc("/movies/{id}", handlers.GetMoviesId).Methods("GET")
	s.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	s.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")
	s.HandleFunc("/moviesComputation/", handlers.MovieComp).Methods("GET")
	// Bind to a port and pass our router in
	logger.WriteLogInfo("main", "Start Server on port 8000", "start server")
	log.Fatal(http.ListenAndServe(":8000", r))
}
