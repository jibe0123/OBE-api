package main

import (
	v1 "api/internal/app/router/v1"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()

	// Route group for API v1
	apiV1 := r.Group("/api/v1")
	v1.RegisterRoutes(apiV1) // Registering API v1 routes

	// Server configuration and start
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Starting the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("Server running on http://localhost:8080 🚀")

	// Waiting for interrupt signal for a graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Gracefully shutting down the server... 🛑")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Forced server shutdown: ", err)
	}
	fmt.Println("Server stopped 🛑")
}
