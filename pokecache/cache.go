package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mutex: &sync.RWMutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	entry, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
