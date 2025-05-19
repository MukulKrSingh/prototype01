package config

import (
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
	// Load .env file if it exists
	_ = godotenv.Load()

	return &Config{
		Server: ServerConfig{
			Port: getEnv("PORT", defaultPort),
		},
		MongoDB: MongoDBConfig{
			URI:      getEnv("MONGODB_URI", getEnv("MDB_MCP_CONNECTION_STRING", defaultMongoURI)),
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
