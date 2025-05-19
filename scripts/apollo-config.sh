#!/bin/bash
# This script generates an Apollo Studio compatible Config file

# Set variables
GRAPHQL_ENDPOINT="http://localhost:8080/graphql"
OUTPUT_DIR="./apollo"
CONFIG_FILE="$OUTPUT_DIR/apollo.config.js"

echo "📝 Generating Apollo Studio configuration..."

# Create directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Generate the Apollo config file
cat > "$CONFIG_FILE" << EOL
// Apollo Studio configuration file
module.exports = {
  client: {
    service: {
      name: "prototype01",
      url: "$GRAPHQL_ENDPOINT",
    },
    includes: ["./apollo/queries/**/*.{js,ts,graphql}"],
  },
};
EOL

# Create queries directory
mkdir -p "$OUTPUT_DIR/queries"

# Create example query file
cat > "$OUTPUT_DIR/queries/example.graphql" << EOL
# Example GraphQL query for Apollo Studio
query PingExample {
  ping
}
EOL

echo "✅ Apollo Studio configuration created!"
echo "📚 Configuration file: $CONFIG_FILE"
echo ""
echo "To use Apollo Studio with your API:"
echo "1. Open Apollo Studio: https://studio.apollographql.com/sandbox/explorer"
echo "2. Enter endpoint URL: $GRAPHQL_ENDPOINT"
echo "3. Run your queries!"
