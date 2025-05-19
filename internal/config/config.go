package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all application configuration
type Config struct {
	Server  ServerConfig
	MongoDB MongoDBConfig
	Env     string
}

// ServerConfig holds server specific configuration
type ServerConfig struct {
	Port string
}

// MongoDBConfig holds MongoDB specific configuration
type MongoDBConfig struct {
	URI      string
	Database string
}

// Default configuration values
const (
	defaultPort          = "8080"
	defaultMongoURI      = "mongodb://localhost:27017"
	defaultMongoDatabase = "ecommerce"
	defaultEnvironment   = "development"
)

// Load loads configuration from environment variables and .env file
func Load() (*Config, error) {
	// Try loading .env from multiple possible locations
	envPaths := []string{
		".env",                     // Current working directory
		"../.env",                  // One directory up
		"../../.env",               // Two directories up (if running from cmd/server)
		os.Getenv("PWD") + "/.env", // Absolute path from PWD
	}

	var loadedPath string
	var loadErr error

	for _, path := range envPaths {
		log.Printf("Attempting to load .env from: %s\n", path)
		err := godotenv.Load(path)
		if err == nil {
			loadedPath = path
			log.Printf("Successfully loaded .env from: %s\n", path)
			break
		}
		loadErr = err
	}

	if loadedPath == "" {
		log.Printf("Warning: Failed to load .env file: %v\n", loadErr)
		log.Printf("Current working directory: %s\n", getCwd())
	}

	// Log the MongoDB connection string being used
	mongoUri := getEnv("MONGODB_URI", getEnv("MDB_MCP_CONNECTION_STRING", defaultMongoURI))
	log.Println("Using MongoDB URI:", mongoUri)

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", defaultPort),
		},
		MongoDB: MongoDBConfig{
			URI:      mongoUri,
			Database: getEnv("MONGODB_DATABASE", defaultMongoDatabase),
		},
		Env: getEnv("ENV", defaultEnvironment),
	}, nil
}

// Helper to get environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper to get current working directory
func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return "unknown"
	}
	return dir
}
