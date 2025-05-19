#!/bin/bash

# Script for testing the Gin-based API

# Start the server in the background
echo "Starting the server..."
go run cmd/server/main.go &
SERVER_PID=$!

# Wait for the server to start
echo "Waiting for server to start..."
sleep 3

# Test the health endpoint
echo "Testing health endpoint..."
HEALTH_RESPONSE=$(curl -s http://localhost:8080/health)
echo "Health Response: $HEALTH_RESPONSE"

# Test the GraphQL endpoint with a simple query
echo "Testing GraphQL endpoint..."
GRAPHQL_RESPONSE=$(curl -s -X POST \
  -H "Content-Type: application/json" \
  -d '{"query": "{ ping }"}' \
  http://localhost:8080/graphql)
echo "GraphQL Response: $GRAPHQL_RESPONSE"

# Stop the server
echo "Stopping server..."
kill $SERVER_PID

echo "Test complete!"
