package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu      sync.Mutex
	entries map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	return nil
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	ticker := time.NewTicker(interval)

	go func() {
		for t := range ticker.C {
			for key, entry := range c.entries {
				if t.Sub(entry.createdAt).Seconds() > interval.Seconds() {
					delete(c.entries, key)
				}
			}
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		mu:      sync.Mutex{},
		entries: make(map[string]cacheEntry),
	}

	cache.reapLoop(interval)

	return &cache
}
