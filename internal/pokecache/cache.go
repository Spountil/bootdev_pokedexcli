package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	c := Cache{
		CacheMap: map[string]CacheEntry{},
		Interval: interval,
		mu:       sync.RWMutex{},
	}

	go c.reapLoop()

	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	result, ok := c.CacheMap[key]
	if ok {
		return result.Val, ok
	}
	return []byte{}, ok
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.Interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mu.Lock()
		for key, val := range c.CacheMap {
			age := now.Sub(val.CreatedAt)
			if age > c.Interval {
				delete(c.CacheMap, key)
			}
		}
		c.mu.Unlock()
	}
}
