package cache

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func(data []byte) uint32

type ConsistentHashMap struct {
	hash     HashFunc       // hash算法
	replicas int            // 虚拟节点的数量
	keys     []int          // 排序后的key
	hashMap  map[int]string // 内部hashMap
}

func NewConsistentHash(replicas int, hashFunc HashFunc) *ConsistentHashMap {
	consistentHashMap := &ConsistentHashMap{
		replicas: replicas,
		hash:     hashFunc,
		hashMap:  make(map[int]string),
	}
	if consistentHashMap.hash == nil {
		consistentHashMap.hash = crc32.ChecksumIEEE
	}
	return consistentHashMap
}

func (consistentHashMap *ConsistentHashMap) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < consistentHashMap.replicas; i++ {
			hash := int(consistentHashMap.hash([]byte(strconv.Itoa(i) + key)))
			// 放入排序数组
			consistentHashMap.keys = append(consistentHashMap.keys, hash)
			// 存储到hash表
			consistentHashMap.hashMap[hash] = key
		}
	}
	// 内置的排序函数
	sort.Ints(consistentHashMap.keys)
}

func (consistentHashMap *ConsistentHashMap) Get(key string) string {
	if len(consistentHashMap.keys) == 0 {
		return ""
	}

	hash := int(consistentHashMap.hash([]byte(key)))
	// 二分查找
	idx := sort.Search(len(consistentHashMap.keys), func(i int) bool {
		return consistentHashMap.keys[i] >= hash
	})

	return consistentHashMap.hashMap[consistentHashMap.keys[idx%len(consistentHashMap.keys)]]
}
