package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createAt time.Time
	val []byte
}

type Cache struct {
	entries map[string]cacheEntry
	interval time.Duration
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	return &Cache{
		entries: nil,
		interval: interval,
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createAt: time.Now(),
		val: val,
	}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

