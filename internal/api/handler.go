// Package api contains the GraphQL API implementation
package api

import (
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/prototype01/internal/api/cache"
	"github.com/prototype01/internal/api/generated"
	"github.com/prototype01/internal/api/resolvers"
	"github.com/prototype01/internal/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewHandler creates a new GraphQL handler using gqlgen
func NewHandler(db *mongo.Client) (http.Handler, error) {
	// Create a new resolver with the DB
	resolver := &resolvers.Resolver{DB: db}

	// Create a config with the resolver
	config := generated.Config{Resolvers: resolver}

	// Create a new handler with the executable schema
	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	// Add extensions and middleware for Apollo Studio support
	h.Use(extension.Introspection{})

	// Add Apollo Tracing support - using built-in tracing features
	// Create an in-memory cache with a 1 hour TTL
	inMemoryCache := cache.NewInMemoryCache(1 * time.Hour)
	h.Use(extension.AutomaticPersistedQuery{
		Cache: inMemoryCache,
	})

	// Configure WebSocket transport for subscriptions (if needed)
	h.AddTransport(transport.Options{})
	h.AddTransport(transport.GET{})
	h.AddTransport(transport.POST{})
	h.AddTransport(transport.MultipartForm{})

	// Uncomment when needed:
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

// RegisterHandlers registers all API handlers with the Gin router
func RegisterHandlers(router *gin.Engine, client *mongo.Client) error {
	// Create GraphQL handler
	graphqlHandler, err := NewHandler(client)
	if err != nil {
		return err
	}

	// Create playground handler
	playgroundHandler := NewPlaygroundHandler("/graphql")

	// GraphQL endpoint with CORS for Apollo Studio
	router.POST("/graphql", func(c *gin.Context) {
		graphqlHandler.ServeHTTP(c.Writer, c.Request)
	})
	router.OPTIONS("/graphql", middleware.GinCORSMiddleware(), func(c *gin.Context) {
		c.Status(204)
	})

	// GraphQL playground
	router.GET("/playground", func(c *gin.Context) {
		playgroundHandler.ServeHTTP(c.Writer, c.Request)
	})

	return nil
}
