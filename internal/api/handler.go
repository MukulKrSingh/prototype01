// Package api contains the GraphQL API implementation
package api

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/prototype01/internal/api/generated"
	"github.com/prototype01/internal/api/resolvers"
	"go.mongodb.org/mongo-driver/mongo"
	// These will be uncommented in Step 2
	// "github.com/99designs/gqlgen/graphql/handler/extension"
	// "github.com/99designs/gqlgen/graphql/handler/transport"
)

// NewHandler creates a new GraphQL handler using gqlgen
func NewHandler(db *mongo.Client) (http.Handler, error) {
	// Create a new resolver with the DB
	resolver := &resolvers.Resolver{DB: db}

	// Create a config with the resolver
	config := generated.Config{Resolvers: resolver}

	// Create a new handler with the executable schema
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	// Add extensions and middleware
	// These will be properly configured in Step 2
	// h.Use(extension.Introspection{})
	// h.AroundOperations(middleware.OperationMiddleware())
	// h.AroundResponses(middleware.ResponseMiddleware())

	return h, nil
}

// NewPlaygroundHandler creates a GraphQL playground handler
// This provides an interactive UI for testing GraphQL queries
func NewPlaygroundHandler(endpoint string) http.Handler {
	title := "GraphQL Playground - E-commerce API"
	return playground.Handler(title, endpoint)
}
