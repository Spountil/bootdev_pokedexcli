package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	CacheMap map[string]cacheEntry
	Interval time.Duration
	mu       sync.RWMutex
}
