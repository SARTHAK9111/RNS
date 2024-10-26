package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// Starts the Redis subscriber in a Go routine, keeping it active
func StartRedisSubscriber(redisClient *redis.Client) {
	go func() {
		pubsub := redisClient.Subscribe(context.Background(), "input_notifications", "periodic_notifications")
		defer pubsub.Close()

		for {
			// Listen for new messages from Redis
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				continue
			}
			log.Printf("Received new message: %s", msg.Payload)
		}
	}()
}

// Handler to acknowledge the subscriber
func RedisSubscriberHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Listening for notifications...")
	}
}
