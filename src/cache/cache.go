package cache

import (
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *Cache
	cacheBytes int64
}

func (cache *cache) add(key string, value ByteView) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.lru == nil {
		cache.lru = New(cache.cacheBytes, nil)
	}
	cache.lru.Add(key, value)
}

func (cache *cache) get(key string) (value ByteView, ok bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if cache.lru == nil {
		return
	}

	if v, ok := cache.lru.Get(key); ok {
		return v.(ByteView), ok
	}
	return
}
