package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// WSHandler handles WebSocket connections and sends Redis notifications
func WSHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP connection to WebSocket
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			log.Printf("Failed to upgrade connection: %v", err)
			return
		}
		defer conn.Close()

		// Subscribe to the Redis "notifications" channel
		pubsub := redisClient.Subscribe(context.Background(), "notifications")
		defer pubsub.Close()

		for {
			// Wait for a message from Redis
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Printf("Failed to receive Redis message: %v", err)
				return
			}

			// Send the Redis message over WebSocket
			if err := wsutil.WriteServerText(conn, []byte(msg.Payload)); err != nil {
				log.Printf("Failed to send WebSocket message: %v", err)
				return
			}
		}
	}
}
