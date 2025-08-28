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

func (c Cache) Add(key string, val []byte) {
        c.mu.Lock()
        defer c.mu.Unlock()
        c.mapEntries[key] = cacheEntry{createdAT: time.Now(), val: val} 
}

func (c Cache) Get(key string) []byte, bool {
        c.mu.Lock()
        defer c.mu.Unlock()
        
