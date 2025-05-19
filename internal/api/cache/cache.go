// Package cache provides caching implementations for GraphQL
package cache

import (
	"context"
	"sync"
	"time"
)

// InMemoryCache provides a simple in-memory cache implementation for APQ
type InMemoryCache struct {
	cache map[string]string
	mutex sync.RWMutex
	ttl   time.Duration
}

// NewInMemoryCache creates a new in-memory cache with the given TTL
func NewInMemoryCache(ttl time.Duration) *InMemoryCache {
	return &InMemoryCache{
		cache: make(map[string]string),
		ttl:   ttl,
	}
}

// Add adds a value to the cache
func (c *InMemoryCache) Add(ctx context.Context, key string, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = value
	// In a real implementation, we'd handle TTL by starting a cleanup goroutine
}

// Get retrieves a value from the cache
func (c *InMemoryCache) Get(ctx context.Context, key string) (string, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	val, found := c.cache[key]
	return val, found
}
