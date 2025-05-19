#!/bin/bash
# This script starts the server in development mode with GraphQL playground

echo "🚀 Starting development server with GraphQL playground..."

# Set environment variables
export ENV=development
export PORT=8080

# Run the server
go run main.go

# This script will exit when the server is stopped
