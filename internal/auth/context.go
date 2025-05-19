package auth

import (
	"context"
	"net/http"
)

// RequestContextKey is the key used to store the HTTP request in the context
const RequestContextKey contextKey = "httpRequest"

// GetRequestFromContext retrieves the HTTP request from the context
func GetRequestFromContext(ctx context.Context) *http.Request {
	if reqValue := ctx.Value(RequestContextKey); reqValue != nil {
		if req, ok := reqValue.(*http.Request); ok {
			return req
		}
	}
	return nil
}
