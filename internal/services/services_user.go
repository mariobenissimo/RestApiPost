package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mariobenissimo/RestApiPost/internal/models"
	"github.com/mariobenissimo/RestApiPost/pkg/db"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) *models.User {
	//INSERT INTO "user" (id, email, password) VALUES ('f47b4840-775a-4f85-babb-cdf746fd2502', 'marioben' , 'ciao')
	insertStatement := `INSERT INTO "user" (id, email, password) VALUES ($1, $2, $3)`

	// Prepare the statement
	stmt, err := db.DB.Prepare(insertStatement)
	if err != nil {
		panic(err)

	}
	defer stmt.Close()

	// Execute the statement with the values
	_, err = stmt.Exec(user.Id, user.Email, user.Password)
	if err != nil {
		panic(err)
	}
	return &user
}
func FindOne(user models.User) map[string]interface{} {
	query := `SELECT * FROM "user" WHERE email = $1 LIMIT 1`
	row := db.DB.QueryRow(query, user.Email)

	// Retrieve the result into a User struct
	var loginUser models.User
	err := row.Scan(&loginUser.Id, &loginUser.Email, &loginUser.Password)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := models.Token{
		Id:    loginUser.Id,
		Email: loginUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = loginUser
	return resp
}
