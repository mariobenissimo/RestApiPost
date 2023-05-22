package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func InizializeDatabase() {
	// Connessione al database
	connStr := "host=localhost port=5432 user=mario password=example dbname=mario sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// defer db.Close()

	// Test della connessione
	err = DB.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connessione al database riuscita!")
}
