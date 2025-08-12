package lru

import (
	"fmt"
	"sync"
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Eviction", func(t *testing.T) {
		cache := New(2)
		cache.Put("a", 1)
		cache.Put("b", 2)
		cache.Put("c", 3)

		if _, ok := cache.Get("a"); ok {
			t.Error("key 'a' should have been evicted")
		}

		if _, ok := cache.Get("b"); !ok {
			t.Error("key 'b' should not have been evicted")
		}

		if _, ok := cache.Get("c"); !ok {
			t.Error("key 'c' should not have been evicted")
		}
	})

	t.Run("Hit", func(t *testing.T) {
		cache := New(2)
		cache.Put("a", 1)
		cache.Put("b", 2)
		cache.Get("a")
		cache.Put("c", 3)

		if _, ok := cache.Get("b"); ok {
			t.Error("key 'b' should have been evicted")
		}

		if _, ok := cache.Get("a"); !ok {
			t.Error("key 'a' should not have been evicted")
		}

		if _, ok := cache.Get("c"); !ok {
			t.Error("key 'c' should not have been evicted")
		}
	})

	t.Run("Concurrency", func(t *testing.T) {
		cache := New(100)
		var wg sync.WaitGroup

		// Concurrent writes
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				cache.Put(fmt.Sprintf("key%d", i), i)
			}(i)
		}
		wg.Wait()

		// Concurrent reads
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				val, ok := cache.Get(fmt.Sprintf("key%d", i))
				if !ok {
					t.Errorf("key%d not found", i)
				}
				if val.(int) != i {
					t.Errorf("expected value %d for key%d, got %d", i, i, val.(int))
				}
			}(i)
		}
		wg.Wait()
	})
}
