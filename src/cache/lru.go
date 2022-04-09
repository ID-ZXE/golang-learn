package cache

import "container/list"

type Cache struct {
	maxBytes    int64
	actualBytes int64
	ll          *list.List
	cache       map[string]*list.Element
	OnEvicted   func(key string, value EntryValue) // optional and executed when an entry is purged.
}

type entry struct {
	key   string
	value EntryValue
}

// EntryValue use Len to count how many bytes it takes
type EntryValue interface {
	Len() int
}

// New is the Constructor of Cache
func New(maxBytes int64, onEvicted func(string, EntryValue)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (cache *Cache) Get(key string) (value EntryValue, ok bool) {
	if element, ok := cache.cache[key]; ok {
		cache.ll.MoveToFront(element)
		// 类型转换 interface.()
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return
}

func (cache *Cache) RemoveOldest() {
	ele := cache.ll.Back()
	if ele != nil {
		cache.ll.Remove(ele)
		// 类型转换
		kv := ele.Value.(*entry)
		delete(cache.cache, kv.key)
		cache.actualBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if cache.OnEvicted != nil {
			cache.OnEvicted(kv.key, kv.value)
		}
	}
}

// Add adds a value to the cache.
func (cache *Cache) Add(key string, value EntryValue) {
	if ele, ok := cache.cache[key]; ok {
		cache.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		cache.actualBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := cache.ll.PushFront(&entry{key, value})
		cache.cache[key] = ele
		cache.actualBytes += int64(len(key)) + int64(value.Len())
	}
	// maxBytes为0表示缓存无限制
	for cache.maxBytes != 0 && cache.maxBytes < cache.actualBytes {
		cache.RemoveOldest()
	}
}

// Len the number of cache entries
func (cache *Cache) Len() int {
	return cache.ll.Len()
}
