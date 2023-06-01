package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/internal/services"
	"github.com/mariobenissimo/RestApiPost/pkg/logger"
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	movie := services.GetMovies(ctx, cancel)
	if ctx.Err() != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, ctx.Err().Error(), "Timeout on getMovies")
		json.NewEncoder(w).Encode(models.Response{"Error": "Timeout"})
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
	if ctx.Err() != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, ctx.Err().Error(), "Timeout on getMoviesId")
		json.NewEncoder(w).Encode(models.Response{"Error": "Timeout"})
	} else {
		json.NewEncoder(w).Encode(movie)
	}
}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		cancel()
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, err.Error(), "Error in decode body of request on updatemovie")
		panic(err)
	}
	services.UpdateMovie(movie, ctx, cancel)
	if ctx.Err() != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, ctx.Err().Error(), "Timeout on updateMovies")
		json.NewEncoder(w).Encode(models.Response{"Error": "Timeout"})
	} else {
		json.NewEncoder(w).Encode(models.Response{"Info": "Record modificato con successo"})
	}
}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	recordCancelled := services.DeleteMovie(id, ctx, cancel)
	if ctx.Err() != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, ctx.Err().Error(), "Timeout on updateMovies")
		json.NewEncoder(w).Encode(models.Response{"Error": "Timeout"})
	}
	if recordCancelled {
		json.NewEncoder(w).Encode(models.Response{"Info": "Record cancellato con successo"})
	} else {
		json.NewEncoder(w).Encode(models.Response{"Info": "Nessun record"})
	}
}
func MovieComp(w http.ResponseWriter, r *http.Request) {

}
func Create(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, err.Error(), "Error on decode body in request on create")
		panic(err)
	}
	services.InsertMovie(movie.Id, movie.Title, movie.Year, movie.Plot, movie.ImdbRating, ctx, cancel)
	services.InsertDirector(movie.Director.Id, movie.Director.Name, movie.Director.Surname, movie.Id, ctx, cancel)
	actors := movie.Actors
	for _, actor := range actors {
		services.InsertActor(actor.Id, actor.Name, actor.Surname, movie.Id, ctx, cancel)
	}
	if ctx.Err() != nil {
		logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, ctx.Err().Error(), "Timeout on updateMovies")
		json.NewEncoder(w).Encode(models.Response{"Error": "Timeout"})
	} else {
		json.NewEncoder(w).Encode(models.Response{"Info": "Record inserito con successo"})
	}
}
