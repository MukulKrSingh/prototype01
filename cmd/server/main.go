// Package server contains the main application entry point
package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
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

	// Set Gin mode based on environment
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create a new Gin router
	router := gin.New()

	// Add middleware
	router.Use(middleware.GinLoggerMiddleware())
	router.Use(middleware.GinRecoveryMiddleware())
	router.Use(middleware.GinContextMiddleware())

	// Enhanced health check endpoint
	router.GET("/health", func(c *gin.Context) {
		// Check MongoDB connection
		mongoStatus := "ok"
		if err := client.Database("admin").RunCommand(c, map[string]interface{}{"ping": 1}).Err(); err != nil {
			mongoStatus = "error: " + err.Error()
		}

		// Return comprehensive health status
		c.JSON(http.StatusOK, gin.H{
			"status":    "running",
			"timestamp": time.Now().Format(time.RFC3339),
			"services": gin.H{
				"mongodb": mongoStatus,
				"api":     "ok",
			},
			"message": "E-Commerce Backend Setup Complete!",
		})
	})

	// Register GraphQL handlers
	if err := api.RegisterHandlers(router, client); err != nil {
		logger.Fatal("Failed to register GraphQL handlers", err)
	}

	// Log availability
	if cfg.Env == "development" {
		logger.Info("GraphQL Playground available at http://localhost:" + cfg.Server.Port + "/playground")
		logger.Info("Apollo Studio can connect to http://localhost:" + cfg.Server.Port + "/graphql")
	}

	// Home page redirects to playground in development mode
	router.GET("/", func(c *gin.Context) {
		if cfg.Env == "development" {
			c.Redirect(http.StatusFound, "/playground")
		} else {
			c.String(200, "E-Commerce Backend API")
		}
	})

	// Handle 404
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not Found")
	})

	// Start server
	port := cfg.Server.Port
	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
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
