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

## Advanced Features

### Automatic Persisted Queries (APQ)

Our GraphQL API supports Automatic Persisted Queries for improved performance. Here's how it works:

#### Client Setup (Apollo Client)
```javascript
import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client';
import { createPersistedQueryLink } from '@apollo/client/link/persisted-queries';
import { sha256 } from 'crypto-hash';

const httpLink = createHttpLink({ 
  uri: 'http://localhost:8080/graphql' 
});

// Set up persisted queries
const persistedQueriesLink = createPersistedQueryLink({ 
  sha256,
  useGETForHashedQueries: true // Use GET for better caching
});

const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: persistedQueriesLink.concat(httpLink),
});
```

#### How APQ Works

1. **First Request**: Client sends a persisted query request with a query ID (SHA-256 hash of the query)
2. **Cache Miss**: If the server doesn't recognize the ID, it responds with an error
3. **Full Query**: Client sends the full query text along with the ID
4. **Cache Storage**: Server stores the query text with its ID in the cache
5. **Subsequent Requests**: Only the query ID is sent, saving bandwidth

#### Benefits
- Reduces request size (especially for large queries)
- Improves network performance
- Reduces server parsing overhead
- Works seamlessly with Apollo Client

These examples will be expanded as more functionality is added in Step 2.
