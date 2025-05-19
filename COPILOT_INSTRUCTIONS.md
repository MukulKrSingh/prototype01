# E-commerce Backend Development Instructions

This document outlines the step-by-step process to build a robust Go-based E-commerce backend with GraphQL API.

## Development Roadmap

### 1. Project Setup (Foundation)
- [x] Initialize Go module
- [x] Set up GVM (Go Version Manager) for version management
- [x] Create proper directory structure following Go standards
- [x] Set up MongoDB connection through MCP
- [x] Add essential Go packages and dependencies

### 2. Core Components Development
- [ ] Design database schema for e-commerce entities
- [ ] Create GraphQL schema definitions
- [ ] Implement GraphQL resolvers and mutations
- [ ] Set up authentication and authorization
- [ ] Create product catalog management
- [ ] Implement order processing system
- [ ] Develop user management functionality

### 3. Code Implementation Details
- [ ] Create domain models for entities (users, products, orders, etc.)
- [ ] Implement repository layer for database operations
- [ ] Develop service layer for business logic
- [ ] Create GraphQL API handlers
- [ ] Implement middleware for security, logging, etc.
- [ ] Set up code generation for GraphQL schema

### 4. Testing and Quality Assurance
- [ ] Write unit tests for crucial components
- [ ] Implement integration tests for API endpoints
- [ ] Set up CI/CD pipeline for automated testing
- [ ] Perform security audits and optimization

### 5. Documentation
- [ ] Create comprehensive README.md with setup instructions
- [ ] Document API endpoints and usage examples
- [ ] Create PROJECT.md for Go concepts explanation
- [ ] Maintain CHANGELOG.md for version history
- [ ] Add code comments for better understanding

## Specific Implementation Approach

1. **Directory Structure**: Follow standard Go project layout with separation of concerns
   - `cmd/`: Application entry points
   - `internal/`: Private application code
   - `pkg/`: Public libraries that can be imported by other projects
   - `api/`: GraphQL schema and resolvers

2. **GraphQL Implementation**:
   - Use `gqlgen` for code generation
   - Define schemas first, then generate resolvers
   - Implement custom resolvers for complex business logic

3. **Database Access**:
   - Use MongoDB with proper repositories
   - Use MongoDB Atlas connection string available in .env file (MDB_MCP_CONNECTION_STRING)
   - Implement interfaces for better testability
   - Use context for proper request handling

4. **Authentication**:
   - JWT-based authentication
   - Role-based access control
   - Middleware for securing endpoints

5. **Documentation**:
   - Keep code comments up-to-date
   - Document all public APIs
   - Provide examples in README

## Best Practices To Follow

1. Use dependency injection for better testability
2. Follow separation of concerns (layers: API, Service, Repository)
3. Handle errors properly with meaningful messages
4. Use context for request scoping and cancellation
5. Follow Go coding conventions (variable naming, error handling, etc.)
6. Use interfaces for better abstraction and testing
7. Keep packages small and focused on single responsibility
8. Use proper logging for debugging and monitoring
9. Implement graceful shutdown for the server
10. Use environment variables for configuration
11. Use the env package to access any launch-specific configuration
