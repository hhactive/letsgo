package lru

import (
	"container/list"
	"sync"
)

// Cache is the interface for a LRU cache.
type Cache interface {
	// Put adds a value to the cache.
	Put(key, value interface{})
	// Get retrieves a value from the cache.
	Get(key interface{}) (interface{}, bool)
}

type lruCache struct {
	capacity int
	mu       sync.Mutex
	ll       *list.List
	cache    map[interface{}]*list.Element
}

type entry struct {
	key   interface{}
	value interface{}
}

// New creates a new LRU cache.
func New(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[interface{}]*list.Element),
	}
}

func (c *lruCache) Put(key, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	elem := c.ll.PushFront(&entry{key, value})
	c.cache[key] = elem

	if c.capacity != 0 && c.ll.Len() > c.capacity {
		last := c.ll.Back()
		if last != nil {
			delete(c.cache, last.Value.(*entry).key)
			c.ll.Remove(last)
		}
	}
}

func (c *lruCache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if elem, ok := c.cache[key]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}

	return nil, false
}