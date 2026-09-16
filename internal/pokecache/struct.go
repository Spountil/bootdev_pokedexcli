package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

type Cache struct {
	CacheMap map[string]CacheEntry
	Interval time.Duration
	mu       sync.RWMutex
}
