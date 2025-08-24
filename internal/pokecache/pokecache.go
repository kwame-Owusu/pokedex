package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte // raw data that we are caching
}
type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry := cacheEntry{createdAt: time.Now(), value: val}
	cache.entries[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	val, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return val.value, true
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for {
		<-ticker.C
		currTime := time.Now()
		cache.mutex.Lock()
		for key, val := range cache.entries {
			if currTime.Sub(val.createdAt) > interval {
				delete(cache.entries, key)
			}
		}
		cache.mutex.Unlock()
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop(interval)
	return cache
}
