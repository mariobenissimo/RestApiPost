package models

import "github.com/google/uuid"

type Actor struct {
	Id      uuid.UUID `db:"idactor" json:"Id"`
	Name    string    `db:"name" json:"Name"`
	Surname string    `db:"surname" json:"Surname"`
}
