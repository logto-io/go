package client

// CacheStorage defines the interface for cache storage
type CacheStorage interface {
	// Get retrieves the cached value
	// if you need set prefix, please implement it in the implementation
	Get(key string) (string, error)
	// Setex sets a cache with expiration time
	Setex(key string, value string, seconds int) error
}