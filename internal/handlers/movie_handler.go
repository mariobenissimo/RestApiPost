package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/internal/services"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	movie := services.GetMovies(ctx, cancel)
	w.Header().Set("Content-Type", "application/json")
	if ctx.Err() != nil {
		json.NewEncoder(w).Encode(`{"error":"timeout"}`)
	} else {
		json.NewEncoder(w).Encode(movie)
	}
}

func GetMoviesId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	movie := services.GetMovieById(id, ctx, cancel)
	w.Header().Set("Content-Type", "application/json")
	if ctx.Err() != nil {
		json.NewEncoder(w).Encode(`{"error":"timeout"}`)
	} else {
		json.NewEncoder(w).Encode(movie)
	}
}
func Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		panic(err)
	}
	services.InsertMovie(movie.Id, movie.Title, movie.Year, movie.Plot, movie.ImdbRating, ctx, cancel)
	services.InsertDirector(movie.Director.Id, movie.Director.Name, movie.Director.Surname, movie.Id, ctx, cancel)
	actors := movie.Actors
	for _, actor := range actors {
		services.InsertActor(actor.Id, actor.Name, actor.Surname, movie.Id, ctx, cancel)
	}
	w.Header().Set("Content-Type", "application/json")
	if ctx.Err() != nil {
		json.NewEncoder(w).Encode(`{"error":"timeout"}`)
	} else {
		json.NewEncoder(w).Encode(`{"data":"movie inserito con successo"}`)
	}
}
