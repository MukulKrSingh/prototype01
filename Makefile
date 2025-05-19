# Go E-commerce Backend Makefile

# Variables
BINARY_NAME=ecommerce-server
MAIN_PATH=./cmd/server
BUILD_DIR=bin
GO=go
GOFMT=gofmt
GQLGEN=github.com/99designs/gqlgen
CONFIG_FILE=gqlgen.yml

.PHONY: all build clean run run-bin test fmt lint generate help deps dev dev-live apollo apollo-studio

# Default target
all: clean fmt generate test build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	@$(GO) build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Build successful! Binary is in $(BUILD_DIR) directory."

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean successful!"

# Run the application
run:
	@echo "Running $(BINARY_NAME)..."
	@$(GO) run $(MAIN_PATH)

# Run the application with the binary
run-bin: build
	@echo "Running $(BINARY_NAME) from binary..."
	@./$(BUILD_DIR)/$(BINARY_NAME)

# Run tests
test:
	@echo "Running tests..."
	@$(GO) test ./... -v

# Format code
fmt:
	@echo "Formatting code..."
	@$(GOFMT) -w .

# Lint code
lint:
	@echo "Linting code..."
	@$(GO) vet ./...
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not found, skipping additional linting"; \
	fi

# Generate GraphQL code
generate:
	@echo "Generating GraphQL code..."
	@if ! command -v gqlgen > /dev/null; then \
		echo "Installing gqlgen..."; \
		$(GO) get $(GQLGEN); \
	fi
	@$(GO) run $(GQLGEN) generate
	@echo "Code generation complete!"

# Start GraphQL development server with playground
dev:
	@echo "Starting GraphQL development server..."
	@ENV=development PORT=8080 $(GO) run $(MAIN_PATH)
	
# Start GraphQL development server with playground and hot reload (requires air)
dev-live:
	@echo "Starting GraphQL development server with hot reload..."
	@if ! command -v air > /dev/null; then \
		echo "Installing air..."; \
		$(GO) install github.com/cosmtrek/air@latest; \
	fi
	@ENV=development PORT=8080 air -c .air.toml

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@$(GO) mod tidy
	@$(GO) mod download
	@echo "Dependencies installed!"

# Setup Apollo Studio config
apollo:
	@echo "Setting up Apollo Studio configuration..."
	@chmod +x ./scripts/apollo-config.sh
	@./scripts/apollo-config.sh
	@echo "Apollo Studio configuration complete!"

# Start server and open Apollo Studio
apollo-studio:
	@echo "Starting server with Apollo Studio integration..."
	@chmod +x ./scripts/apollo-studio.sh
	@./scripts/apollo-studio.sh

# Help command
help:
	@echo "Available commands:"
	@echo "  make build      - Build the application"
	@echo "  make clean      - Clean build artifacts"
	@echo "  make run        - Run the application directly with go run"
	@echo "  make run-bin    - Build and run the application binary"
	@echo "  make dev        - Run with GraphQL Playground for development"
	@echo "  make dev-live   - Run with hot reload using Air"
	@echo "  make test       - Run tests"
	@echo "  make fmt        - Format code"
	@echo "  make lint       - Lint code"
	@echo "  make generate   - Generate GraphQL code using gqlgen"
	@echo "  make apollo     - Generate Apollo Studio configuration"
	@echo "  make apollo-studio - Start server and open Apollo Studio"
	@echo "  make deps       - Install dependencies"
	@echo "  make all        - Clean, format, generate, test, and build"
	@echo "  make help       - Show this help message"
