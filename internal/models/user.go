package models

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `db:"id" json:"Id,omitempty"`
	Email    string    `db:"email" json:"Email"`
	Password string    `db:"password" json:"Password"`
}
type Token struct {
	Id    uuid.UUID `json:"Id"`
	Email string    `json:"Email"`
	jwt.StandardClaims
}
