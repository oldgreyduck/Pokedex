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
        newCache :=  &Cache{
                mapEntries:   make(map[string]cacheEntry)
                interval:     interval
        }
        go newCache.reapLoop(interval)
        return newCache
}

func (c Cache) Add(key string, val []byte) {
        c.mu.Lock()
        defer c.mu.Unlock()
        c.mapEntries[key] = cacheEntry{createdAT: time.Now(), val: val} 
}

func (c Cache) Get(key string) []byte, bool {
        c.mu.Lock()
        defer c.mu.Unlock()
        val, ok := c.mapEntries[key]
        if !ok {
                return nil, false
        }
        return val.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
        ticker := time.NewTicker(interval)
        defer ticker.Stop()
        for {
                select {
                case <- ticker.C:
                        c.mu.Lock()
                        defer c.mu.Unlock()
                        now := time.Now()
                        for key, entry := range c.mapEntries {
                                if now.Sub(entry.createdAt) > interval {
                                        delete(c.mapEntries, key)
                                }
                        }
                }
        }
}

