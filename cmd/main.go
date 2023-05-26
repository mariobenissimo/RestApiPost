package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/mariobenissimo/RestApiPost/internal/handlers"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	log "github.com/sirupsen/logrus"
)

type Middleware struct {
	handler http.Handler
}
type favContextKey string

func (l *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Gestione del middleware tramite log su file
	// in questa fase è possibile creare un contesto della richiesta e salvare ddati utili in tale richiesta (come ad esempio l'istanza dello user utilizzata)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	// log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	log.WithFields(log.Fields{
		"method": r.Method,
		"path":   r.URL.Path,
		"time":   time.Since(start),
	}).Info("Request")
}
func middleware(handlerToWrap http.Handler) *Middleware {
	return &Middleware{handlerToWrap}
}
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		var auth = r.Header.Get("Authorization") //Grab the token from the header
		if auth == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Missing auth token"})
			log.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"error":  "Missing auth token",
			}).Warn("Request login")
			return
		}
		header := strings.Split(auth, "Bearer ")[1]
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"Message": "Missing auth token"})
			log.WithFields(log.Fields{
				"method": r.Method,
				"path":   r.URL.Path,
				"error":  "Missing auth token",
			}).Warn("Request login")
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
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
		panic(err)
	}
	defer f.Close()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(f)
	log.SetLevel(log.InfoLevel)

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
