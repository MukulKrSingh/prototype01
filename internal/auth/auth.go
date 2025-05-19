// Package auth provides authentication utilities for the application
package auth

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/prototype01/pkg/logger"
)

// Key type for context values
type contextKey string

// UserIDKey is the key used to store the user ID in the context
const UserIDKey contextKey = "userID"

// ExtractTokenFromContext extracts the JWT token from the request context
func ExtractTokenFromContext(ctx context.Context) string {
	// Get the request from the context
	request := getRequestFromContext(ctx)
	if request == nil {
		return ""
	}

	// Extract token from Authorization header
	authHeader := request.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// The header format should be "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		logger.Warn("Invalid authorization header format")
		return ""
	}

	return parts[1]
}

// VerifyToken verifies the JWT token and returns the user ID if valid
func VerifyToken(token string) (string, error) {
	// This is a simplified implementation for prototype purposes
	// In a real application, you would use a JWT library to verify the token

	if token == "" {
		return "", errors.New("empty token provided")
	}

	// For the prototype, we'll just check that the token has a minimum length
	// and return a mock user ID
	if len(token) < 10 {
		return "", errors.New("token too short")
	}

	// In a real implementation, you would decode and verify the JWT token here
	// For now, we'll just return a mock user ID
	return "user-123", nil
}

// GetUserIDFromContext retrieves the user ID from the context if present
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

// getRequestFromContext extracts the http.Request from the context
func getRequestFromContext(ctx context.Context) *http.Request {
	// The HTTP request is commonly stored in context by HTTP middleware
	// This might need to be adjusted based on how your GraphQL server is set up
	if requestData := ctx.Value(graphqlRequestContextKey()); requestData != nil {
		if request, ok := requestData.(*http.Request); ok {
			return request
		}
	}
	return nil
}

// graphqlRequestContextKey returns the key used by your GraphQL framework to store the request
func graphqlRequestContextKey() interface{} {
	// This may need to be adjusted based on the GraphQL framework you're using
	// For gqlgen, it's often something like this:
	return contextKey("request")
}
