package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
)

func insertMovie(id uuid.UUID, title string, year string, plot string, imdbRating float32) uuid.UUID {
	// Define the INSERT statement
	insertStatement := `INSERT INTO movie (idmovie, title, year, plot, imdbrating) VALUES ($1, $2, $3, $4, $5)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, title, year, plot, strconv.FormatFloat(float64(imdbRating), 'f', -1, 32))
	if err != nil {
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
	return id
}
func insertActor(id uuid.UUID, name string, surname string, fkmovie uuid.UUID) {
	// Define the INSERT statement
	insertStatement := `INSERT INTO actor (idactor, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
}
func insertDirector(id uuid.UUID, name string, surname string, fkmovie uuid.UUID) {
	// Define the INSERT statement
	insertStatement := `INSERT INTO director (iddirector, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	var rows *sql.Rows
	var err error
	rows, err = db.DB.Query("SELECT * FROM movie")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var movies []models.Movie
	// Itera sui risultati della query
	for rows.Next() {
		var movie models.Movie
		// Scansiona i valori delle colonne nella struttura
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.ImdbRating)
		if err != nil {
			panic(err)
		}
		movie.Actors = getActors(movie.Id)
		movie.Director = getDirector(movie.Id)
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
func getDirector(id uuid.UUID) models.Director {
	var director models.Director
	query := "SELECT iddirector,name,surname FROM director where fkmovie = $1 LIMIT 1"
	err := db.DB.QueryRow(query, id).Scan(&director.Id, &director.Name, &director.Surname)
	if err != nil {
		panic(err)
	}
	return director
}
func getActors(id uuid.UUID) []models.Actor {
	query := "SELECT idactor,name,surname FROM actor where fkmovie = $1 "
	rows, err := db.DB.Query(query, id)
	if err != nil {
		panic(err)
	}
	var actors []models.Actor
	for rows.Next() {
		//fetch actors
		var actor models.Actor
		err = rows.Scan(&actor.Id, &actor.Name, &actor.Surname)
		if err != nil {
			panic(err)
		}
		actors = append(actors, actor)
	}
	return actors
}
func GetMovieById(id string) models.Movie {
	var movie models.Movie
	query := "SELECT * FROM movie where idmovie = $1 LIMIT 1"
	err := db.DB.QueryRow(query, id).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.ImdbRating)
	if err != nil {
		panic(err)
	}
	movie.Actors = getActors(movie.Id)
	movie.Director = getDirector(movie.Id)
	return movie
}

func GetMoviesId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	movie := GetMovieById(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}
func Create(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		panic(err)
	}
	insertMovie(movie.Id, movie.Title, movie.Year, movie.Plot, movie.ImdbRating)
	insertDirector(movie.Director.Id, movie.Director.Name, movie.Director.Surname, movie.Id)

	actors := movie.Actors
	for _, actor := range actors {
		insertActor(actor.Id, actor.Name, actor.Surname, movie.Id)
	}
	Body := `{"data":"movie inserito con successo"}`
	w.Write([]byte(Body))
}
