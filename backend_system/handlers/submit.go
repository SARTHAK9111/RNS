package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

func SubmitHandler(requestContext context.Context, db *sql.DB, redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		content := r.FormValue("content")
		if content == "" {
			http.Error(w, "Content is required", http.StatusBadRequest)
			return
		}

		// Insert content into the database (blocking)
		_, err := db.ExecContext(requestContext, "INSERT INTO submissions (content) VALUES (?)", content)
		if err != nil {
			log.Printf("Database error: %v", err)
			http.Error(w, "Failed to insert content", http.StatusInternalServerError)
			return
		}

		// Publish to Redis in a Go routine (non-blocking)
		go func() {
			err := redisClient.Publish(requestContext, "input_notifications", content).Err()
			if err != nil {
				log.Printf("Redis publish error: %v", err)
			} else {
				log.Printf("Published new message: %s", content)
			}
		}()

		fmt.Fprintln(w, "Input notification \n", content)
	}
}
