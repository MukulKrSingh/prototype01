#!/bin/bash
# This script generates GraphQL code using gqlgen

echo "📝 Generating GraphQL code with gqlgen..."

# Check if gqlgen is installed
if ! command -v gqlgen &> /dev/null; then
    echo "⚙️ Installing gqlgen..."
    go install github.com/99designs/gqlgen@latest
fi

# Generate the code
go run github.com/99designs/gqlgen generate

echo "✅ GraphQL code generation completed!"
echo "📚 See the generated files in internal/api/generated/"
