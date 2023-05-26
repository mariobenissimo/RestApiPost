package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type ErrorResponse struct {
	Err string
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}
	user.Password = string(pass)
	createdUser := services.CreateUser(user)
	if createdUser == nil {
		err := ErrorResponse{
			Err: "Password Encryption failed",
		}
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(createdUser)
}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := services.FindOne(user)
	json.NewEncoder(w).Encode(resp)
}
