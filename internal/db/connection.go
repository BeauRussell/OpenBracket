package db

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AddConfigPath("./.config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read the configuration file: %v", err)
	}

	url := viper.GetString("DB_URL") + "?authToken=" + viper.GetString("DB_AUTH_TOKEN")
	DB, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database is not reachable: %v", err)
	}

	return DB
}
