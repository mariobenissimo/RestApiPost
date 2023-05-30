package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/logger"
)

type favContextKey string

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header)
		var auth = r.Header.Get("Authorization") //Grab the token from the header
		if auth == "" {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Response{"Message": "Missing auth token"})
			logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, "Missing auth token", "Request login error missing auth token")
			return
		}
		header := strings.Split(auth, "Bearer ")[1]
		header = strings.TrimSpace(header)
		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Response{"Message": "Missing auth token"})
			logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, "Missing auth token", "Request login error missing auth token")
			return
		}
		tk := &models.Token{}

		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(models.Response{"Message": err.Error()})
			logger.WriteLogRequesWarnWithError(r.Method, r.URL.Path, "Invalid auth token", "Request login invalid token")
			return
		}
		k := favContextKey("user")
		ctx := context.WithValue(r.Context(), k, tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
