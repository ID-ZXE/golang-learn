package cache

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get look ups a key's value
func (cache *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := cache.cache[key]; ok {
		cache.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest removes the oldest item
func (cache *Cache) RemoveOldest() {
	ele := cache.ll.Back()
	if ele != nil {
		cache.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(cache.cache, kv.key)
		cache.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if cache.OnEvicted != nil {
			cache.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (cache *Cache) Add(key string, value Value) {
	if ele, ok := cache.cache[key]; ok {
		cache.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		cache.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := cache.ll.PushFront(&entry{key, value})
		cache.cache[key] = ele
		cache.nbytes += int64(len(key)) + int64(value.Len())
	}
	for cache.maxBytes != 0 && cache.maxBytes < cache.nbytes {
		cache.RemoveOldest()
	}
}

// Len the number of cache entries
func (cache *Cache) Len() int {
	return cache.ll.Len()
}
