package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prototype01/pkg/logger"
)

// GinContextKey is the key for storing Gin's context in the request context
const GinContextKey = "GinContextKey"

// GinContextMiddleware adds Gin context to the request context
func GinContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// GetGinContext extracts the Gin context from a request context
func GetGinContext(ctx context.Context) (*gin.Context, bool) {
	ginContext, ok := ctx.Value(GinContextKey).(*gin.Context)
	return ginContext, ok
}

// GinLoggerMiddleware logs HTTP requests using our custom logger
func GinLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Log after request is processed
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// Format log message
		logMsg := fmt.Sprintf(
			"Request processed - Method: %s, Path: %s, Status: %d, Latency: %s",
			method,
			path,
			statusCode,
			latency.String(),
		)

		// Log the request
		logger.Info(logMsg)
	}
}

// GinRecoveryMiddleware recovers from panics and logs the error
func GinRecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Format the error message with additional context
				path := c.Request.URL.Path
				method := c.Request.Method
				clientIP := c.ClientIP()
				errMsg := "Panic recovered in HTTP handler: " + method + " " + path + " from " + clientIP

				// Convert panic value to error
				var err error
				switch v := r.(type) {
				case error:
					err = v
				case string:
					err = panicError{v}
				default:
					err = panicError{"unknown panic"}
				}

				// Log the error with our logger
				logger.Error(errMsg, err)

				// Return an error response
				c.AbortWithStatusJSON(500, gin.H{
					"error": "Internal Server Error",
				})
			}
		}()
		c.Next()
	}
}

// panicError wraps a panic value as an error
type panicError struct {
	value interface{}
}

func (e panicError) Error() string {
	if v, ok := e.value.(string); ok {
		return "panic: " + v
	}
	return "panic occurred"
}

// GinCORSMiddleware adds CORS headers for Gin
func GinCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
