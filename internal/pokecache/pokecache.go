package pokecache

import (
        "fmt"
        "time"
        "net/http"
        "sync"
)

type Cache struct {
        mapEntries    map[string]cacheEntry
        mu            sync.Mutex
        interval      time.Duration
}

type cacheEntry struct {
        createdAt     time.Time
        val           []byte
}

func NewCache(interval time.Duration) *Cache {
        return &Cache{
                mapEntries:   make(map[string]cacheEntry)
                interval:     interval
        }
}

func (c Cache) cacheAdd(key string, val []byte) {
        createdAT:    time.Now()
        val:          val
}
