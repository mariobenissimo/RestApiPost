package models

import "github.com/google/uuid"

type Director struct {
	Id      uuid.UUID `db:"iddirector" json:"Id"`
	Name    string    `db:"name" json:"Name"`
	Surname string    `db:"surname" json:"Surname"`
}
