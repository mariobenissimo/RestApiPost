package testpost

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	"github.com/mariobenissimo/RestApiPost/pkg/middleware"
	"github.com/stretchr/testify/assert"
)

func getToken(router *mux.Router, t *testing.T) string {
	payload := `{"email": "mariobenissimo@gmail.com", "password": "ciao"}`
	router.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	req, err := http.NewRequest("POST", "/login", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	var response models.ResponseLogin
	// get token from request
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
		return ""
	}
	return response.Token
}
func createMovie(router *mux.Router, t *testing.T, token string) models.Response {
	router.HandleFunc("/movies", handlers.DeleteMovie)
	req, err := http.NewRequest("POST", "/movies", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	//log.Println(rr.Body)
	var response models.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
		return nil
	}
	return response
}
func TestDeleteMovie(t *testing.T) {
	db.InizializeDatabase()
	r := mux.NewRouter()
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(middleware.JwtVerify)
	token := getToken(r, t)
	response := createMovie(s, t, token)
	assert.Equal(t, response, models.Response{"Info": "Record cancellato con successo"})
}
