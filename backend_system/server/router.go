package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"Realtime-Notification-System/backend_system/handlers"
	middleware "Realtime-Notification-System/backend_system/middleware_layer"

	"github.com/go-redis/redis/v8"
)

// Router sets up the HTTP server with Redis and graceful shutdown
func Router(db *sql.DB) error {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Use WSL IP if needed
	})

	// Create a request-scoped context for handlers
	requestContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Set up the /submit route with CORS middleware
	mux.Handle("/submit", middleware.EnableCORS(
		handlers.SubmitHandler(requestContext, db, redisClient),
	))

	// CORS for /notifications route is commented out because we run the Redis
	// subscription as a background Go routine, making the route redundant.
	// Uncomment if you need to expose notifications via HTTP for testing or fallback purposes.
	/*
		mux.Handle("/notifications", middleware.EnableCORS(
			handlers.RedisSubscriber(requestContext, redisClient),
		))
	*/

	// Create and start the HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start Redis Subscriber in a Go routine
	log.Println("Enabled Redis Input Subscriber")
	handlers.StartRedisSubscriber(redisClient)

	log.Println("Enabled Redis Periodic Subscriber")
	go handlers.StartClockPublisher(redisClient, db)

	// Run the server in a separate Go routine
	go func() {
		log.Println("Server running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Websocket
	log.Println("Web Socket Established")
	mux.HandleFunc("/ws", handlers.WSHandler(redisClient))

	// Handle OS signals for graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop // Wait for termination signal
	log.Println("Shutting down server...")

	// Create a shutdown context with a timeout
	shutdownContext, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := server.Shutdown(shutdownContext); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	// Close Redis client connection
	if err := redisClient.Close(); err != nil {
		log.Fatalf("Failed to close Redis client: %v", err)
	}

	log.Println("Server stopped gracefully")
	return nil
}
