// Package middlewares provides GraphQL-specific middleware functions
package middlewares

import (
	"context"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/prototype01/internal/api/resolvers"
	"github.com/prototype01/internal/auth"
	"github.com/prototype01/pkg/logger"
)

// OperationMiddleware logs GraphQL operations and adds timing information
func OperationMiddleware() graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// Record start time
		startTime := time.Now()

		// Get operation name and type from the context
		op := graphql.GetOperationContext(ctx)
		operationType := op.Operation.Operation
		operationName := op.OperationName

		// Log the operation start
		logger.Info("GraphQL operation started: " + string(operationType) + " " + operationName)

		// Process the request
		resp := next(ctx)

		// Calculate duration
		duration := time.Since(startTime)

		// Log the operation completion
		logger.Info("GraphQL operation completed: " + string(operationType) + " " + operationName + " in " + duration.String())

		return resp
	}
}

// ResponseMiddleware adds timing information to GraphQL responses
func ResponseMiddleware() graphql.ResponseMiddleware {
	return func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		// Get the response
		resp := next(ctx)

		// Add server timing header in extensions if not present
		if resp.Extensions == nil {
			resp.Extensions = make(map[string]interface{})
		}

		// Get operation info
		op := graphql.GetOperationContext(ctx)

		// Add timing information to extensions
		resp.Extensions["timing"] = map[string]interface{}{
			"parsing":    op.Stats.Parsing.End.Sub(op.Stats.Parsing.Start).Milliseconds(),
			"validation": op.Stats.Validation.End.Sub(op.Stats.Validation.Start).Milliseconds(),
			"read":       op.Stats.Read.End.Sub(op.Stats.Read.Start).Milliseconds(),
			"total":      time.Since(op.Stats.OperationStart).Milliseconds(),
		}

		return resp
	}
}

// AuthMiddleware handles authentication for GraphQL operations
func AuthMiddleware(resolver *resolvers.Resolver) graphql.OperationMiddleware {
	return func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// Get HTTP request from context using graphqlRequestContextKey
		// Since transport.GetHTTP is not available in this version of gqlgen
		// We'll use our own custom auth package to get the request
		httpRequest := auth.GetRequestFromContext(ctx)
		if httpRequest == nil {
			return next(ctx)
		}

		// Store the request in context for later token extraction (this is redundant but keeping for clarity)
		ctx = context.WithValue(ctx, auth.RequestContextKey, httpRequest)

		// Extract token from request
		token := extractToken(httpRequest)

		if token != "" {
			// Verify the token
			userID, err := auth.VerifyToken(token)
			if err == nil {
				// If token is valid, set user in context
				ctx = context.WithValue(ctx, auth.UserIDKey, userID)
				logger.Info("Authenticated user: " + userID)
			} else {
				logger.Warn("Invalid authentication token: " + err.Error())
			}
		}

		// Continue with the operation regardless of auth status
		// Individual resolvers can check auth as needed
		return next(ctx)
	}
}

// extractToken gets the JWT token from the Authorization header
func extractToken(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return ""
	}

	// The header format should be "Bearer <token>"
	const prefix = "Bearer "
	if len(bearerToken) > len(prefix) && bearerToken[:len(prefix)] == prefix {
		return bearerToken[len(prefix):]
	}

	return ""
}
