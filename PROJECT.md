# Go E-Commerce Project Understanding Guide

This document aims to help you understand the Go language concepts used in this project, especially if you're new to Go.

## Table of Contents

- [Go Language Basics](#go-language-basics)
- [Project Structure Explained](#project-structure-explained)
- [Key Go Concepts Used](#key-go-concepts-used)
- [GraphQL Implementation](#graphql-implementation)
- [MongoDB Integration](#mongodb-integration)
- [Common Patterns Used](#common-patterns-used)

## Go Language Basics

### What is Go?

Go (or Golang) is a statically typed, compiled language designed at Google. It's known for its simplicity, efficiency, and built-in concurrency features.

### Key Go Features

1. **Simplicity**: Go has a clean syntax with only ~25 keywords
2. **Concurrency**: Goroutines and channels make concurrent programming easier
3. **Compiled**: Go compiles to a single binary with no external dependencies
4. **Garbage Collection**: Automatic memory management
5. **Strong Typing**: Type safety with compile-time checking

### Go vs Other Languages

If you're coming from other languages:

- **vs JavaScript/Node.js**: Go is compiled, statically typed, and has no callback hell
- **vs Python**: Go is much faster and compiles to a binary
- **vs Java**: Go has simpler syntax, faster compilation, and built-in concurrency

## Project Structure Explained

Our project follows the standard Go project layout:

### `/cmd`

Contains the main applications for the project. Each application has its own directory.

```go
// Example from cmd/server/main.go
package main

import (
    "log"
    "net/http"
    
    "github.com/your/project/internal/api"
    "github.com/your/project/internal/config"
)

func main() {
    cfg := config.Load()
    server := api.NewServer(cfg)
    log.Fatal(http.ListenAndServe(":"+cfg.Port, server))
}
```

### `/internal`

Contains private application code that shouldn't be imported by other projects.

```go
// Example from internal/domain/model/product.go
package model

type Product struct {
    ID          string  `json:"id" bson:"_id,omitempty"`
    Name        string  `json:"name" bson:"name"`
    Description string  `json:"description" bson:"description"`
    Price       float64 `json:"price" bson:"price"`
    CategoryID  string  `json:"categoryId" bson:"category_id"`
}
```

### `/pkg`

Contains code that can be used by external applications.

```go
// Example from pkg/logger/logger.go
package logger

import "log"

// Info logs an informational message
func Info(message string) {
    log.Printf("[INFO] %s", message)
}

// Error logs an error message
func Error(err error) {
    log.Printf("[ERROR] %s", err)
}
```

## Key Go Concepts Used

### 1. Interfaces

Interfaces in Go define behavior. They allow for loose coupling between components.

```go
// Repository interface
type ProductRepository interface {
    FindByID(ctx context.Context, id string) (*model.Product, error)
    FindAll(ctx context.Context) ([]*model.Product, error)
    Create(ctx context.Context, product *model.Product) error
    Update(ctx context.Context, product *model.Product) error
    Delete(ctx context.Context, id string) error
}
```

### 2. Structs and Methods

Structs are collections of fields, and methods can be added to structs:

```go
// Product service
type ProductService struct {
    repo ProductRepository
}

// Method on ProductService
func (s *ProductService) GetProduct(ctx context.Context, id string) (*model.Product, error) {
    return s.repo.FindByID(ctx, id)
}
```

### 3. Error Handling

Go uses explicit error handling:

```go
func GetUserByID(id string) (*User, error) {
    if id == "" {
        return nil, errors.New("id cannot be empty")
    }
    
    user, err := repository.FindUserByID(id)
    if err != nil {
        return nil, fmt.Errorf("finding user: %w", err)
    }
    
    return user, nil
}
```

### 4. Context

Context is used for cancellation signals, deadlines, and request-scoped values:

```go
func (r *productResolver) Products(ctx context.Context) ([]*model.Product, error) {
    // The ctx can be used to check if the client has disconnected
    // or to pass request-scoped values
    return r.service.GetAllProducts(ctx)
}
```

## GraphQL Implementation

We use [gqlgen](https://gqlgen.com/) for our GraphQL implementation, which generates code from our schema.

### Schema Definition

```graphql
# api/schema.graphql
type Product {
  id: ID!
  name: String!
  description: String
  price: Float!
  category: Category
}

type Query {
  products: [Product!]!
  product(id: ID!): Product
}

type Mutation {
  createProduct(input: CreateProductInput!): Product!
}
```

### Resolver Example

```go
// Generated resolver with custom implementation
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
    return r.ProductService.GetProduct(ctx, id)
}
```

## MongoDB Integration

We use the official MongoDB Go driver for database operations:

```go
// Example MongoDB repository implementation
func (r *productRepository) FindByID(ctx context.Context, id string) (*model.Product, error) {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
    
    var product model.Product
    err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&product)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("product not found")
        }
        return nil, err
    }
    
    return &product, nil
}
```

## Common Patterns Used

### 1. Dependency Injection

We use constructor-based dependency injection:

```go
// Service depends on repository
func NewProductService(repo ProductRepository) *ProductService {
    return &ProductService{repo: repo}
}

// Handler depends on service
func NewProductHandler(service *ProductService) *ProductHandler {
    return &ProductHandler{service: service}
}
```

### 2. Repository Pattern

We abstract database operations behind interfaces:

```go
// Interface
type UserRepository interface {
    FindByID(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
}

// Implementation
type MongoUserRepository struct {
    collection *mongo.Collection
}

func (r *MongoUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    // Implementation details...
}
```

### 3. Middleware Pattern

Middleware functions modify or enhance the behavior of HTTP handlers:

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        
        // Validate token...
        
        // Call the next handler
        next.ServeHTTP(w, r)
    })
}
```

### 4. Service Layer

The service layer implements business logic and uses repositories for data access:

```go
type OrderService struct {
    orderRepo  OrderRepository
    productRepo ProductRepository
    userRepo   UserRepository
}

func (s *OrderService) CreateOrder(ctx context.Context, orderInput *CreateOrderInput) (*Order, error) {
    // Validate products exist
    // Check inventory
    // Calculate totals
    // Create order
    return s.orderRepo.Save(ctx, newOrder)
}
```

This guide should help you understand the key Go concepts used in the project. As you work with the codebase, refer back to this document whenever you encounter unfamiliar patterns or concepts.
