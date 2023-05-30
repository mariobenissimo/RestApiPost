package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/mariobenissimo/RestApiPost/pkg/logger"
)

var DB *sql.DB

func InizializeDatabase() {
	connStr := "host=localhost port=5432 user=mario password=example dbname=mario sslmode=disable"
	var err error

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.WriteLogError("InizializeDatabase", err.Error(), "Error opening Database")
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		logger.WriteLogError("InizializeDatabase", err.Error(), "Error contacting Database")
		panic(err)
	}
	logger.WriteLogInfo("InizializeDatabase", "Connessione al database riuscita!", "Database info")
}
