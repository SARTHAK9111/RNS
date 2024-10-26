package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// StartClockPublisher runs a non-blocking periodic Redis publisher
func StartClockPublisher(redisClient *redis.Client, db *sql.DB) {
	ticker := time.NewTicker(15 * time.Second) // GRC: Creates a ticker every minute
	defer ticker.Stop()                        // GRC: Ensures the ticker stops properly

	for t := range ticker.C {
		message := fmt.Sprintf("Event Driven Notification %s", t.Format("15:04:05"))
		err := redisClient.Publish(context.Background(), "notifications", message).Err()
		if err != nil {
			log.Printf("Failed to publish message: %v", err)
		} else {
			log.Printf("Published: %s", message)
		}

		// Insert into event_notifications table
		_, err = db.Exec("INSERT INTO events (event) VALUES (?)",
			message)
		if err != nil {
			log.Printf("Failed to insert into events: %v", err)
		} /*else {
			log.Printf("Message stored in DB: %s", message)
		}*/ // Not needed to log each time
	}
}
