package db

import (
	"database/sql"
	"fmt"
)

func inizializedatabase() {
	// Connessione al database
	connStr := "host=localhost port=5432 user=mario password=example dbname=mario sslmode=disable"
	var err error

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// Test della connessione
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connessione al database riuscita!")
}
