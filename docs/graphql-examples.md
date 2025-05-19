# GraphQL API Examples

This document provides examples of GraphQL queries and mutations for the e-commerce API.

## Basic Queries

### Health Check
```graphql
{
  ping
}
```

Response:
```json
{
  "data": {
    "ping": "GraphQL Server is running! Current time: 2025-05-19T10:30:00Z"
  }
}
```

### Version Information
```graphql
{
  version {
    number
    buildDate
    environment
  }
}
```

Response:
```json
{
  "data": {
    "version": {
      "number": "0.1.0",
      "buildDate": "2025-05-19T10:30:00Z",
      "environment": "development"
    }
  }
}
```

## Working with gqlgen

### Development Workflow

1. **Define your schema**: Add or modify types in the `api/graphql/*.graphql` files
2. **Generate code**: Run `make generate` to create Go types and resolver interfaces
3. **Implement resolvers**: Fill in the resolver implementations in `internal/api/resolvers/`
4. **Test**: Use the GraphQL playground at `http://localhost:8080/playground` during development

### Debugging

When debugging gqlgen generated code:

1. Check the resolver implementation to ensure it matches the schema
2. Use the GraphQL playground to test queries in isolation
3. Add logging to your resolvers to trace execution
4. Check the context for user authentication information

### Common Patterns

#### Pagination
```graphql
{
  products(first: 10, after: "cursor") {
    edges {
      node {
        id
        name
      }
      cursor
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
}
```

#### Filtering
```graphql
{
  products(filter: { category: "electronics", minPrice: 100 }) {
    id
    name
    price
  }
}
```

#### Authentication
```graphql
mutation {
  login(email: "user@example.com", password: "password123") {
    token
    user {
      id
      name
    }
  }
}
```

These examples will be expanded as more functionality is added in Step 2.
