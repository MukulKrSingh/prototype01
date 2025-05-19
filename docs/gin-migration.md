# Migration to Gin Framework

This document describes the migration of our GraphQL API from the standard Go HTTP package to the Gin framework.

## Why Gin?

Gin is a web framework for Go that provides:

1. **Performance**: Gin is built for speed with a martini-like API
2. **Middleware**: Rich middleware ecosystem and easy-to-use middleware interface
3. **Routing**: Fast and flexible HTTP router
4. **Binding**: Easy request data binding (JSON, XML, Form, etc.)
5. **Validation**: Built-in request validation
6. **Error Handling**: Simplified error management
7. **Response Rendering**: JSON, XML, HTML, etc.

## Migration Changes

The following components were updated during the migration:

### 1. Middleware

- Created Gin-specific middleware in `internal/middleware/gin_middleware.go`
- Implemented context-passing middleware to ensure GraphQL operations have access to Gin's context
- Added Gin-compatible CORS middleware for Apollo Studio

### 2. Context Management

- Enhanced `auth` package to support both standard HTTP and Gin context models
- Added functions to extract authentication tokens from Gin context
- Created context key hierarchy that maintains backward compatibility

### 3. Server Configuration

- Updated `main.go` to use Gin's router instead of standard HTTP
- Configured Gin mode based on environment (development/production)
- Added proper error handling and graceful shutdown

### 4. Request Processing

- Modified GraphQL handler registration to work with Gin
- Updated playground handler to work with Gin
- Added support for GraphQL over HTTP GET, POST, and multipart

## Using Gin Features

Now that we use Gin, you can take advantage of its features:

### Request Binding

```go
type LoginInput struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

func LoginHandler(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    // Process login...
}
```

### Route Groups

```go
// API v1 routes
v1 := router.Group("/api/v1")
{
    v1.GET("/products", listProducts)
    v1.GET("/products/:id", getProduct)
    v1.POST("/products", createProduct)
}
```

### File Uploads

```go
router.POST("/upload", func(c *gin.Context) {
    file, _ := c.FormFile("file")
    // Save file...
    c.String(200, "File uploaded successfully")
})
```

## Testing with Gin

Gin provides testing utilities that make writing HTTP tests easier:

```go
func TestLoginAPI(t *testing.T) {
    router := setupRouter()
    
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"email":"test@example.com","password":"password"}`))
    req.Header.Set("Content-Type", "application/json")
    
    router.ServeHTTP(w, req)
    
    assert.Equal(t, 200, w.Code)
    // Assert response body...
}
```
