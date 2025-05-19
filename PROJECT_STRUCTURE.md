# Project Structure Documentation

This document provides a detailed overview of the project's folder structure and explains the responsibilities of each component.

## Overall Architecture

The project follows a clean architecture approach with clear separation of concerns:

```
prototype01/
├── api/            # GraphQL schema definitions
├── apollo/         # Apollo Studio configurations and example queries
├── cmd/            # Main application entry points
├── docs/           # Project documentation
├── internal/       # Internal packages (not importable by external projects)
├── logs/           # Application logs
├── pkg/            # Reusable packages that could be imported by other projects
└── scripts/        # Utility scripts for development and deployment
```

## Top-Level Folders

### `/api`

Contains API definitions, primarily GraphQL schema files.

- `/api/graphql/` - GraphQL schema definitions (*.graphql files)
  - `/api/graphql/schema.graphql` - Main GraphQL schema definition
- `/api/schema/` - Additional schema-related files
  - `/api/schema/schema.graphql` - Schema file used by the GraphQL generator
- `/api/README.md` - Documentation specific to the API structure

### `/apollo`

Configuration and examples for Apollo Studio integration.

- `/apollo/queries/` - Example GraphQL queries and mutations for testing
  - `/apollo/queries/auth.graphql` - Authentication-related queries and mutations
  - `/apollo/queries/basic.graphql` - Basic API operations
  - `/apollo/queries/example_queries.graphql` - Example query operations following naming conventions
  - `/apollo/queries/example_mutations.graphql` - Example mutation operations following naming conventions
- `/apollo/apollo.config.js` - Apollo client configuration

### `/cmd`

Main application entry points. Each subdirectory corresponds to a separate executable.

- `/cmd/server/` - Main backend server application

### `/docs`

Project documentation files.

- `/docs/graphql.md` - GraphQL implementation guide with gqlgen
- `/docs/graphql-examples.md` - Examples of GraphQL operations and common patterns
- `/docs/graphql-naming-conventions.md` - Naming standards for queries/mutations
- `/docs/apollo-studio.md` - Guide for using Apollo Studio with this API

### `/internal`

Internal packages that are specific to this project and not meant to be imported by external applications.

- `/internal/api/` - GraphQL API implementation
  - `/internal/api/generated/` - Auto-generated GraphQL code (models.go, generated.go)
  - `/internal/api/handler.go` - GraphQL handler setup and configuration
  - `/internal/api/middlewares/` - GraphQL-specific middleware
  - `/internal/api/resolvers/` - GraphQL resolver implementations (resolver.go, resolvers.go)

- `/internal/auth/` - Authentication and authorization logic
  - `/internal/auth/auth.go` - Authentication utilities
  - `/internal/auth/context.go` - Context-related authentication helpers

- `/internal/config/` - Application configuration
  - `/internal/config/config.go` - Configuration loading and validation

- `/internal/domain/` - Core business domain
  - `/internal/domain/models/` - Domain models/entities
    - `/internal/domain/models/base.go` - Shared model structures and interfaces
    - `/internal/domain/models/models.go` - E-commerce domain entities (User, Product, Order, etc.)
  - `/internal/domain/services/` - Domain business logic services
    - `/internal/domain/services/services.go` - Service interfaces and implementations

- `/internal/middleware/` - HTTP middleware components
  - `/internal/middleware/cors.go` - CORS middleware for Apollo Studio compatibility
  - `/internal/middleware/middleware.go` - General middleware implementations including logging and recovery

- `/internal/repository/` - Data access layer
  - `/internal/repository/mongodb/` - MongoDB implementations

- `/internal/service/` - Application services layer

### `/logs`

Contains application logs.

- `/logs/app/` - Main application logs
- `/logs/mcp/` - Model Context Protocol related logs

### `/pkg`

Library code that can be used by external applications.

- `/pkg/logger/` - Logging utilities
- `/pkg/utils/` - General utility functions
- `/pkg/validator/` - Data validation utilities

### `/scripts`

Utility scripts for various development and operational tasks.

- `/scripts/apollo-config.sh` - Apollo Studio configuration generator
- `/scripts/apollo-config.sh` - Apollo Studio configuration generator
- `/scripts/apollo-studio.sh` - Script to start server with Apollo Studio
- `/scripts/dev-server.sh` - Development server startup script
- `/scripts/generate-graphql.sh` - GraphQL code generation script

## Key Files

- `gqlgen.yml` - GraphQL code generator configuration
- `Makefile` - Project build and management commands
- `go.mod` & `go.sum` - Go module dependencies
- `README.md` - Project overview and quick start guide
- `PROJECT.md` - Detailed project understanding guide
- `SETUP_SUMMARY.md` - Summary of completed setup steps
- `CHANGELOG.md` - Version history and release notes
- `PROJECT_STRUCTURE.md` - This document, explaining the project structure

### GraphQL-Related Files

- `api/graphql/schema.graphql` - Main GraphQL schema definition
- `internal/api/generated/generated.go` - Auto-generated GraphQL server code
- `internal/api/generated/models.go` - Auto-generated GraphQL data models
- `internal/api/handler.go` - GraphQL server configuration
- `internal/api/resolvers/resolver.go` - GraphQL resolver root
- `internal/api/resolvers/resolvers.go` - GraphQL resolver implementations
- `apollo/queries/*.graphql` - Example GraphQL operations

## Architecture Principles

This project follows these architectural principles:

1. **Clean Architecture** - Clear separation between layers with dependencies pointing inward
2. **Domain-Driven Design** - Business logic organized around domain models
3. **Dependency Injection** - Components receive their dependencies rather than creating them
4. **Hexagonal Architecture** - Core business logic is isolated from external concerns

## Flow of Data

The typical flow of data in a request is:

1. HTTP request comes into the server
2. Middleware processes the request (logging, CORS, recovery, authentication)
3. GraphQL handler parses and validates the operation
4. GraphQL resolvers execute the operation
5. Resolvers call domain services
6. Domain services implement business logic
7. Repositories handle data access
8. Response flows back through the layers
9. Middleware processes the response
10. Response returns to the client

This structure ensures separation of concerns, testability, and maintainability as the project grows.
