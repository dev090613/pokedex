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
