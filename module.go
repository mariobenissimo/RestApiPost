package main

import "github.com/google/uuid"

type Movie struct {
	Id         uuid.UUID `db:"idmoview" json:"Id"`
	Title      string    `db:"title" json:"Title"`
	Year       string    `db:"year" json:"Year"`
	Director   Director  `json:"director"`
	Actors     []Actor   `json:"actors"`
	Plot       string    `db:"plot" json:"Plot"`
	ImdbRating float32   `db:"imdbrating" json:"ImdbRating"`
}

type Director struct {
	Id      uuid.UUID `db:"iddirector" json:"Id"`
	Name    string    `db:"name" json:"Name"`
	Surname string    `db:"surname" json:"Surname"`
}
type Actor struct {
	Id      uuid.UUID `db:"idactor" json:"Id"`
	Name    string    `db:"name" json:"Name"`
	Surname string    `db:"surname" json:"Surname"`
}
