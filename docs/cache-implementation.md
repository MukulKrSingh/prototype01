# Cache Implementation for GraphQL

This document covers the cache implementations used in the E-Commerce Backend API.

## Automatic Persisted Query (APQ) Cache

Automatic Persisted Queries (APQ) is a technique to improve network performance for GraphQL applications. Instead of sending the full query text for every request, clients can send a shorter hash of the query to reduce request size.

### Implementation

We've implemented an in-memory cache to support APQ:

1. **Location**: `internal/api/cache/cache.go`
2. **Cache Type**: In-memory with mutex-based synchronization
3. **Default TTL**: 1 hour

### InMemoryCache

The `InMemoryCache` struct provides a simple in-memory cache implementation:

```go
type InMemoryCache struct {
    cache map[string]string
    mutex sync.RWMutex
    ttl   time.Duration
}
```

Key features:
- Thread-safe with read-write mutex
- Configurable TTL
- Simple API with `Add` and `Get` methods

### Usage in GraphQL Handler

The cache is integrated with gqlgen's Automatic Persisted Query extension:

```go
inMemoryCache := cache.NewInMemoryCache(1 * time.Hour)
h.Use(extension.AutomaticPersistedQuery{
    Cache: inMemoryCache,
})
```

### Future Improvements

For production environments, consider implementing:

1. **MongoDB-backed cache**: Replace the in-memory implementation with a MongoDB collection
2. **Redis-backed cache**: Use Redis for distributed caching if needed
3. **Actual TTL implementation**: Add goroutine-based cleanup for expired entries
4. **Metrics and monitoring**: Add instrumentation to track cache hit/miss rates

## Other Caching Opportunities

Besides APQ, consider implementing caching for:

1. **Database query results**: Cache frequently accessed data
2. **Authentication tokens**: JWT token validation results
3. **Product catalog data**: For read-heavy e-commerce operations

## Implementation Notes

When implementing APQ with gqlgen, always ensure:

1. The cache implementation satisfies the required interface
2. A cache is provided when enabling AutomaticPersistedQuery
3. The cache has proper error handling for edge cases