package services

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/google/uuid"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	"github.com/mariobenissimo/RestApiPost/pkg/logger"
)

func InsertMovie(id uuid.UUID, title string, year string, plot string, imdbRating float32, ctx context.Context, cancel context.CancelFunc) uuid.UUID {
	insertStatement := `INSERT INTO movie (idmovie, title, year, plot, imdbrating) VALUES ($1, $2, $3, $4, $5)`
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		logger.WriteLogError("InserMovie", err.Error(), "Error with statment insert movie")
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, title, year, plot, strconv.FormatFloat(float64(imdbRating), 'f', -1, 32))
	if err != nil {
		cancel()
		logger.WriteLogError("InserMovie", err.Error(), "Error with excute insert movie")
		panic(err)
	}
	logger.WriteLogInfo("InserMovie", "Record insert successfully", "Record inserted")
	return id
}
func InsertActor(id uuid.UUID, name string, surname string, fkmovie uuid.UUID, ctx context.Context, cancel context.CancelFunc) {
	insertStatement := `INSERT INTO actor (idactor, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		logger.WriteLogError("InsertActor", err.Error(), "Error with statment insert actor")
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		cancel()
		logger.WriteLogError("InsertActor", err.Error(), "Error with statment insert actor")
		panic(err)
	}
	logger.WriteLogInfo("InsertActor", "Record insert successfully", "Record inserted")

}
func InsertDirector(id uuid.UUID, name string, surname string, fkmovie uuid.UUID, ctx context.Context, cancel context.CancelFunc) {
	insertStatement := `INSERT INTO director (iddirector, name, surname, fkmovie) VALUES ($1, $2, $3, $4)`
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		cancel()
		logger.WriteLogError("InsertDirector", err.Error(), "Error with statment insert director")
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, name, surname, fkmovie)
	if err != nil {
		cancel()
		logger.WriteLogError("InsertDirector", err.Error(), "Error with statment insert director")
		panic(err)
	}
	logger.WriteLogInfo("InserDirector", "Record insert successfully", "Record inserted")
}

func GetDirector(id uuid.UUID, ctx context.Context, cancel context.CancelFunc) models.Director {
	var director models.Director
	query := "SELECT iddirector,name,surname FROM director where fkmovie = $1 LIMIT 1"
	err := db.DB.QueryRow(query, id).Scan(&director.Id, &director.Name, &director.Surname)
	if err != nil {
		cancel()
		logger.WriteLogError("GetDirector", err.Error(), "Error with statment get director")
		panic(err)
	}
	return director
}
func GetActors(id uuid.UUID, ctx context.Context, cancel context.CancelFunc) []models.Actor {
	query := "SELECT idactor,name,surname FROM actor where fkmovie = $1 "
	rows, err := db.DB.Query(query, id)
	if err != nil {
		cancel()
		logger.WriteLogError("GetActors", err.Error(), "Error with statment get actors")
		panic(err)
	}
	var actors []models.Actor
	defer rows.Close()
	for rows.Next() {
		var actor models.Actor
		err = rows.Scan(&actor.Id, &actor.Name, &actor.Surname)
		if err != nil {
			cancel()
			logger.WriteLogError("GetActor", err.Error(), "Error with statment get actor")
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
		logger.WriteLogError("GetMovieById", err.Error(), "Error with statment get moviebyid")
		panic(err)
	}
	movie.Actors = GetActors(movie.Id, ctx, cancel)
	movie.Director = GetDirector(movie.Id, ctx, cancel)
	return movie
}

func GetMovies(ctx context.Context, cancel context.CancelFunc) []models.Movie {
	var rows *sql.Rows
	var err error
	rows, err = db.DB.Query("SELECT * FROM movie")
	if err != nil {
		cancel()
		logger.WriteLogError("GetMovies", err.Error(), "Error with statment getmovies")
		panic(err)
	}
	defer rows.Close()
	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Year, &movie.Plot, &movie.ImdbRating)
		if err != nil {
			logger.WriteLogError("GetMovies", err.Error(), "Error with scan rows")
			panic(err)
		}
		movie.Actors = GetActors(movie.Id, ctx, cancel)
		movie.Director = GetDirector(movie.Id, ctx, cancel)
		movies = append(movies, movie)
	}
	return movies
}
func UpdateMovie(movie models.Movie, ctx context.Context, cancel context.CancelFunc) {
	updateStatement := ` UPDATE movie SET title = $1, year= $2 plot=$3 imdbrating= $4 WHERE idmovie = $5`
	stmt, err := db.DB.Prepare(updateStatement)
	if err != nil {
		logger.WriteLogError("UpdateMovie", err.Error(), "Error with statment updatemovie")
		cancel()
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(movie.Title, movie.Year, movie.Plot, movie.ImdbRating, movie.Id)
	if err != nil {
		logger.WriteLogError("UpdateMovie", err.Error(), "Error with exec updatemovie")
		cancel()
		panic(err)
	}
	updateDirector(movie.Id, movie.Director, ctx, cancel)
	updateActors(movie.Id, movie.Actors, ctx, cancel)
	logger.WriteLogInfo("UpdateMovie", "Record update successfully!", "record update")
}
func updateDirector(id uuid.UUID, director models.Director, ctx context.Context, cancel context.CancelFunc) {
	updateStatement := ` UPDATE director SET name = $1, surname= $2 fkmovie=$3 WHERE iddirector = $4`
	stmt, err := db.DB.Prepare(updateStatement)
	if err != nil {
		logger.WriteLogError("updateDirector", err.Error(), "Error with exec updateDirector")
		cancel()
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(director.Name, director.Surname, id, director.Id)
	if err != nil {
		logger.WriteLogError("updateDirector", err.Error(), "Error with exec updateDirector")
		cancel()
		panic(err)
	}
	logger.WriteLogInfo("updateDirector", "Record update successfully!", "record update")

}
func updateActors(id uuid.UUID, actors []models.Actor, ctx context.Context, cancel context.CancelFunc) {
	for index, actor := range actors {
		updateStatement := ` UPDATE actor SET name = $1, surname= $2 fkmovie=$3 WHERE idactor = $4`
		stmt, err := db.DB.Prepare(updateStatement)
		if err != nil {
			cancel()
			logger.WriteLogError("updateActors", err.Error(), "Error with statment "+strconv.Itoa(index))
			panic(err)
		}
		defer stmt.Close()
		_, err = stmt.Exec(actor.Name, actor.Surname, id, actor.Id)
		if err != nil {
			cancel()
			logger.WriteLogError("updateActors", err.Error(), "Error with exec updateDirector"+strconv.Itoa(index))
			panic(err)
		}
	}
	logger.WriteLogInfo("updateActors", "Record update successfully!", "all record update")

}
func recordExists(id string, ctx context.Context, cancel context.CancelFunc) bool {
	selectQuery := "SELECT * FROM movie WHERE idmovie = $1 LIMIT 1"
	rows, err := db.DB.Query(selectQuery, id)
	if err != nil {
		cancel()
		logger.WriteLogError("recordExists", err.Error(), "Error with statment recordExists")
		panic(err)
	}
	defer rows.Close()
	return rows.Next()
}
func DeleteMovie(id string, ctx context.Context, cancel context.CancelFunc) bool {
	exists := recordExists(id, ctx, cancel)
	if !exists {
		return false
	}
	deleteStatement := ` DELETE from movie WHERE idmovie = $1`
	stmt, err := db.DB.Prepare(deleteStatement)
	if err != nil {
		logger.WriteLogError("DeleteMovie", err.Error(), "Error with statment deleteDirector")
		cancel()
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		cancel()
		logger.WriteLogError("DeleteMovie", err.Error(), "Error with exec deleteDirector")
		panic(err)
	}
	return true
}
func InsertFirstMovie() {
	id := InsertMovie(uuid.New(), "Inception", "2010", "A thief who steals corporate secrets through the use of dream-sharing technology is given the inverse task of planting an idea into the mind of a C.E.O.", 8.8, nil, nil)
	InsertActor(uuid.New(), "Leonardo", "Di Caprio", id, nil, nil)
	InsertActor(uuid.New(), "Joseph", "Gordon-Levitt", id, nil, nil)
	InsertActor(uuid.New(), "Ellen", "Page", id, nil, nil)
	InsertActor(uuid.New(), "Tom", "Hardy", id, nil, nil)
	InsertDirector(uuid.New(), "Christopher", "Nolan", id, nil, nil)
}
