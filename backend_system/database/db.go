package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDB initializes the database with a provided DSN
func InitDB(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}

	log.Println("Database connected successfully")
	return db
}
