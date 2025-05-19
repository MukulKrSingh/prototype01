// Package server contains the main application entry point
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prototype01/internal/api"
	"github.com/prototype01/internal/config"
	"github.com/prototype01/internal/middleware"
	"github.com/prototype01/internal/repository/mongodb"
	"github.com/prototype01/pkg/logger"
)

// main is kept for direct execution of this file
func main() {
	RunServer()
}

// RunServer initializes and starts the HTTP server
func RunServer() {
	// Initialize logger
	logger.Init()
	logger.Info("Starting e-commerce server setup")

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", err)
	}

	// Connect to MongoDB
	client, err := mongodb.Connect(context.Background(), cfg.MongoDB.URI)
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			logger.Error("Failed to disconnect from MongoDB", err)
		}
	}()
	logger.Info("MongoDB connection established")

	// Create a new server mux
	mux := http.NewServeMux()

	// Basic health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "E-Commerce Backend Setup Complete!")
	})

	// GraphQL handler (to be implemented in Step 2)
	graphqlHandler, err := api.NewHandler(client)
	if err != nil {
		logger.Fatal("Failed to create GraphQL handler", err)
	}

	// Set up GraphQL endpoint with CORS for Apollo Studio
	mux.Handle("/graphql", middleware.CORSMiddleware(graphqlHandler))

	// Set up GraphQL playground in development mode
	if cfg.Env == "development" {
		playgroundHandler := api.NewPlaygroundHandler("/graphql")
		mux.Handle("/playground", playgroundHandler)
		logger.Info("GraphQL Playground available at http://localhost:" + cfg.Server.Port + "/playground")
		logger.Info("Apollo Studio can connect to http://localhost:" + cfg.Server.Port + "/graphql")
	}

	// Home page redirects to playground in development mode
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		if cfg.Env == "development" {
			http.Redirect(w, r, "/playground", http.StatusFound)
		} else {
			fmt.Fprintf(w, "E-Commerce Backend API")
		}
	})

	// Apply middleware
	handler := middleware.LoggingMiddleware(middleware.RecoveryMiddleware(mux))

	// Start server
	port := cfg.Server.Port
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Server is running on http://localhost:" + port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP server error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", err)
	}

	logger.Info("Server exited")
}
