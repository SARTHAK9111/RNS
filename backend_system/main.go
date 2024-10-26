package main

import (
	"log"
	"os"

	"Realtime-Notification-System/backend_system/database"
	"Realtime-Notification-System/backend_system/server"

	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the DSN from environment variables
	dsn := os.Getenv("MYSQL_USER") + ":" +
		os.Getenv("MYSQL_PASSWORD") + "@tcp(" +
		os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") +
		")/" + os.Getenv("MYSQL_DB")

	// Initialize the database with the DSN
	db := database.InitDB(dsn)
	defer db.Close()

	//log.Println("Server running on http://localhost:8080")
	log.Fatal(server.Router(db))

	//log.Println("Application started successfully")

}
