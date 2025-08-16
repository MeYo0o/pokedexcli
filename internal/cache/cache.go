package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry    map[string]CacheEntry
	mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	Val       []byte
	CreatedAt time.Time
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entry:    make(map[string]CacheEntry),
		interval: interval,
	}

	go c.reapLoop()

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Entry[key] = CacheEntry{
		Val: val, CreatedAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if val, ok := c.Entry[key]; ok {
		return val.Val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()

		for k, v := range c.Entry {
			if time.Since(v.CreatedAt) > c.interval {
				delete(c.Entry, k)
			}
		}

		c.mu.Unlock()

	}
}
