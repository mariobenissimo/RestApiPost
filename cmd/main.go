package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
)

type Middleware struct {
	handler http.Handler
}
type favContextKey string

func (l *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Gestione del middleware tramite log su file
	// in questa fase è possibile creare un contesto della richiesta e salvare ddati utili in tale richiesta (come ad esempio l'istanza dello user utilizzata)
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

func middleware(handlerToWrap http.Handler) *Middleware {
	return &Middleware{handlerToWrap}
}
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println(r.Header)
		var header = strings.Split(r.Header.Get("Authorization"), " ")[1] //Grab the token from the header
		fmt.Println("QUI" + header)
		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Missing auth token"})
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		//autenticazione del token
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": err.Error()})
			return
		}
		k := favContextKey("user")
		ctx := context.WithValue(r.Context(), k, tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	db.InizializeDatabase()
	//services.InsertFirstMovie()
	r := mux.NewRouter()
	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	// Auth route
	s := r.PathPrefix("/auth").Subrouter()
	s.Use(JwtVerify)
	s.HandleFunc("/movies", handlers.GetMovies).Methods("GET")
	s.HandleFunc("/movies/{id}", handlers.GetMoviesId).Methods("GET")
	s.HandleFunc("/movies", handlers.Create).Methods("POST")
	s.HandleFunc("/movies/{id}", handlers.UpdateMovie).Methods("PUT")
	s.HandleFunc("/movies/{id}", handlers.DeleteMovie).Methods("DELETE")

	wrappedMux := middleware(r)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", wrappedMux))
}
