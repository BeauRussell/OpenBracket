package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() *sql.DB {
	DB, err := sql.Open("sqlite3", "internal/db/OpenBracket.db")
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	return DB
}
