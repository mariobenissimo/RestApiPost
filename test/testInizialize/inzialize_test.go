package testinizialize

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	"github.com/stretchr/testify/assert"
)

func TestInizialize(t *testing.T) {
	// Create a new instance of the router
	router := mux.NewRouter()
	db.InizializeDatabase()
	// Register your routes and handlers
	router.HandleFunc("/getMovies", handlers.GetMovies)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/getMovies", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	var movies []models.Movie
	err = json.Unmarshal(rr.Body.Bytes(), &movies)
	if err != nil {
		// Handle error
		t.Fatal(err)
		return
	}
	assert.Equal(t, len(movies), 1)
}
