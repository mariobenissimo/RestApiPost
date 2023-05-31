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
	router := mux.NewRouter()
	db.InizializeDatabase()
	router.HandleFunc("/getMovies", handlers.GetMovies)
	req, err := http.NewRequest("GET", "/getMovies", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	var movies []models.Movie
	err = json.Unmarshal(rr.Body.Bytes(), &movies)
	if err != nil {
		t.Fatal(err)
		return
	}
	assert.Equal(t, len(movies), 1)
}
