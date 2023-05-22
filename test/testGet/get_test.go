package testget

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestGetMovie(t *testing.T) {
	// Create a new instance of the router
	router := mux.NewRouter()
	// inizializeDatabase()
	// Register your routes and handlers
	router.HandleFunc("/getMovies/{id}", handlers.GetMoviesId)
	url := "/getMovies/9ca5af9a-fba5-4777-acd7-eb39d720dcad"
	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)
	var movie models.Movie
	err = json.Unmarshal(rr.Body.Bytes(), &movie)
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.Equal(t, movie.ImdbRating, float32(8.8))
	var id uuid.UUID
	id, err = uuid.Parse("9ca5af9a-fba5-4777-acd7-eb39d720dcad")
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.Equal(t, movie.Id, id)
}
