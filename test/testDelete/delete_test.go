package testdelete

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
func deleteMovie(router *mux.Router, t *testing.T, token string) models.Response {
	router.HandleFunc("/movies/{id}", handlers.DeleteMovie)
	url := "/auth/movies/9ca5af9a-fba5-4777-acd7-eb39d720dcad"
	req, err := http.NewRequest("DELETE", url, nil)
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
	response := deleteMovie(s, t, token)
	assert.Equal(t, response, models.Response{"Info": "Record cancellato con successo"})
}

// func TestFooerTableDriven(t *testing.T) {
// 	// Defining the columns of the table
// 	var tests = []struct {
// 		name  string
// 		input int
// 		want  string
// 	}{
// 		// the table itself
// 		{"9 should be Foo", 9, "Foo"},
// 		{"3 should be Foo", 3, "Foo"},
// 		{"1 is not Foo", 1, "1"},
// 		{"0 should be Foo", 0, "Foo"},
// 	}
// 	// The execution loop
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ans := Fooer(tt.input)
// 			if ans != tt.want {
// 				t.Errorf("got %s, want %s", ans, tt.want)
// 			}
// 		})
// 	}
// }
