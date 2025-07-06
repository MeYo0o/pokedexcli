package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry    map[string]cacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		Entry:    make(map[string]cacheEntry),
		interval: interval,
	}

	// Start the reap loop in a goroutine
	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entry[key] = cacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if v, ok := c.Entry[key]; ok {
		return v.Val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		for key, entry := range c.Entry {
			if now.Sub(entry.CreatedAt) > c.interval {
				delete(c.Entry, key)
			}
		}
		c.mu.Unlock()
	}
}
