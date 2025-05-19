#!/bin/bash
# This script starts the server and opens Apollo Studio

# Set variables
PORT=8080
GRAPHQL_ENDPOINT="http://localhost:$PORT/graphql"
APOLLO_STUDIO_URL="https://studio.apollographql.com/sandbox/explorer?endpoint=$GRAPHQL_ENDPOINT"

echo "🚀 Starting Go E-Commerce Server with Apollo Studio support..."

# Run server in background
ENV=development PORT=$PORT go run ./cmd/server &
SERVER_PID=$!

echo "⌛ Waiting for server to start..."
sleep 3

# Open Apollo Studio in the default browser
echo "🔗 Opening Apollo Studio Explorer..."
if [[ "$OSTYPE" == "darwin"* ]]; then
  # macOS
  open "$APOLLO_STUDIO_URL"
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
  # Linux
  xdg-open "$APOLLO_STUDIO_URL" &> /dev/null
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
  # Windows
  start "$APOLLO_STUDIO_URL"
else
  echo "Please open Apollo Studio manually at: $APOLLO_STUDIO_URL"
fi

echo "✅ Server is running at http://localhost:$PORT"
echo "📊 Apollo Studio Explorer is connected to your GraphQL API"
echo "Press Ctrl+C to stop the server"

# Handle shutdown
function cleanup {
  echo -e "\n🛑 Shutting down server..."
  kill $SERVER_PID
  echo "Server stopped"
  exit 0
}

trap cleanup INT TERM

# Wait for user to press Ctrl+C
wait $SERVER_PID
