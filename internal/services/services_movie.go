package services

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
)

func InsertMovie(id uuid.UUID, title string, year string, plot string, imdbRating float32, ctx context.Context, cancel context.CancelFunc) uuid.UUID {
	// Define the INSERT statement
	insertStatement := `INSERT INTO movie (idmovie, title, year, plot, imdbrating) VALUES ($1, $2, $3, $4, $5)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, title, year, plot, strconv.FormatFloat(float64(imdbRating), 'f', -1, 32))
	if err != nil {
		cancel()
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
	return id
}
func InsertActor(id uuid.UUID, name string, surname string, fkmovie uuid.UUID, ctx context.Context, cancel context.CancelFunc) {
	// Define the INSERT statement
	insertStatement := `INSERT INTO actor (idactor, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		cancel()
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
}
func InsertDirector(id uuid.UUID, name string, surname string, fkmovie uuid.UUID, ctx context.Context, cancel context.CancelFunc) {
	// Define the INSERT statement
	insertStatement := `INSERT INTO director (iddirector, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		panic(err)
	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		cancel()
		panic(err)
	}

	fmt.Println("Record inserted successfully!")
}

func GetDirector(id uuid.UUID, ctx context.Context, cancel context.CancelFunc) models.Director {
	var director models.Director
	query := "SELECT iddirector,name,surname FROM director where fkmovie = $1 LIMIT 1"
	err := db.DB.QueryRow(query, id).Scan(&director.Id, &director.Name, &director.Surname)
	if err != nil {
		cancel()
		panic(err)
	}
	return director
}
func GetActors(id uuid.UUID, ctx context.Context, cancel context.CancelFunc) []models.Actor {
	query := "SELECT idactor,name,surname FROM actor where fkmovie = $1 "
	rows, err := db.DB.Query(query, id)
	if err != nil {
		cancel()
		panic(err)
	}
	var actors []models.Actor
	for rows.Next() {
		//fetch actors
		var actor models.Actor
		err = rows.Scan(&actor.Id, &actor.Name, &actor.Surname)
		if err != nil {
			cancel()
			panic(err)
		}
		actors = append(actors, actor)
	}
	return actors
}

func GetMovieById(id string, ctx context.Context, cancel context.CancelFunc) models.Movie {
	var movie models.Movie
	query := "SELECT * FROM movie where idmovie = $1 LIMIT 1"
	err := db.DB.QueryRow(query, id).Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.ImdbRating)
	if err != nil {
		cancel()
		panic(err)
	}
	movie.Actors = GetActors(movie.Id, ctx, cancel)
	movie.Director = GetDirector(movie.Id, ctx, cancel)
	return movie
}

func GetMovies(ctx context.Context, cancel context.CancelFunc) []models.Movie {
	// time.Sleep(10 * time.Second)
	var rows *sql.Rows
	var err error
	rows, err = db.DB.Query("SELECT * FROM movie")
	if err != nil {
		cancel()
		fmt.Println(err)
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
		movie.Actors = GetActors(movie.Id, ctx, cancel)
		movie.Director = GetDirector(movie.Id, ctx, cancel)
		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		cancel()
		panic(err)
	}
	return movies
}
