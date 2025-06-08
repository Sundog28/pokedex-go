package pokecache

import "sync"

type Cache interface {
	Get(key string) ([]byte, bool)
	Set(key string, data []byte)
}

type MemoryCache struct {
	mu    sync.RWMutex
	cache map[string][]byte
}

func NewMemoryCache() *MemoryCache {
	return &MemoryCache{cache: make(map[string][]byte)}
}

func (c *MemoryCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.cache[key]
	return data, ok
}

func (c *MemoryCache) Set(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = data
}
