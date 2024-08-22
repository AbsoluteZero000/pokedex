package pokecache

import (
	"time"
	"sync"
)


type CacheEntry struct {
	createdAt time.Time
	data []byte
}

type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.RWMutex
}

func NewCache(duration time.Duration) Cache{
	c := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.RWMutex{},
	}

	go c.reapLoop(duration)

	return c
}


func (c *Cache)Add(key string, data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		createdAt: time.Now(),
		data:      data,
	}

}

func (c *Cache)Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if _, ok := c.cache[key]; !ok {
		return nil, false
	}
	return c.cache[key].data, true
}

func (c *Cache)reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		c.reap(time.Now(), duration)
	}
}

func (c *Cache)reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.cache {
		if now.Sub(v.createdAt) > last {
			delete(c.cache, k)
		}
	}
}
