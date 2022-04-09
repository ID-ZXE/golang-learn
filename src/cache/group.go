package cache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

// 实现了Getter接口的函数对象
type GetterFunc func(key string) ([]byte, error)

type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

func (getterFunc GetterFunc) Get(key string) ([]byte, error) {
	return getterFunc(key)
}

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	group := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	mu.RLock()
	group := groups[name]
	mu.RUnlock()
	return group
}

func (group *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}

	if value, ok := group.mainCache.get(key); ok {
		log.Println("[GeeCache] hit")
		return value, nil
	}

	return group.load(key)
}

func (group *Group) load(key string) (value ByteView, err error) {
	return group.getLocally(key)
}

func (group *Group) getLocally(key string) (ByteView, error) {
	// 加载回调函数
	bytes, err := group.getter.Get(key)
	if err != nil {
		return ByteView{}, err

	}
	value := ByteView{b: cloneBytes(bytes)}
	group.populateCache(key, value)
	return value, nil
}

func (group *Group) populateCache(key string, value ByteView) {
	group.mainCache.add(key, value)
}
