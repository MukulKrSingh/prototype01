package auth

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestContextKey is the key used to store the HTTP request in the context
const RequestContextKey contextKey = "httpRequest"

// UserIDContextKey is the key used to store the authenticated user ID
const UserIDContextKey contextKey = "userId"

// GinContextKey is the key for storing Gin's context in the request context
const GinContextKey = "GinContextKey"

// GetRequestFromContext retrieves the HTTP request from the context
func GetRequestFromContext(ctx context.Context) *http.Request {
	if reqValue := ctx.Value(RequestContextKey); reqValue != nil {
		if req, ok := reqValue.(*http.Request); ok {
			return req
		}
	}
	return nil
}

// GetGinContext extracts the Gin context from a request context
func GetGinContext(ctx context.Context) (*gin.Context, bool) {
	ginContext, ok := ctx.Value(GinContextKey).(*gin.Context)
	return ginContext, ok
}

// GetRequestFromGinContext gets the HTTP request from Gin context
func GetRequestFromGinContext(ctx context.Context) *http.Request {
	ginContext, ok := GetGinContext(ctx)
	if !ok {
		return nil
	}
	return ginContext.Request
}

// GetUserIDFromContext gets the user ID from the context if it exists
func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDContextKey).(string)
	return userID, ok
}
