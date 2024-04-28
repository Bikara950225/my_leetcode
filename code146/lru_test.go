package main

import (
	"testing"
)

func Test_LRUCache(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		cache := ConstructorLRUCache(2)

		cache.Put(1, 1)
		cache.Put(2, 2)
		v1 := cache.Get(1)
		if v1 != 1 {
			t.Errorf("Test_LRUCache(), v1 != 1: %v", v1)
			return
		}
		cache.Put(3, 3)
		v2 := cache.Get(2)
		if v2 != -1 {
			t.Errorf("Test_LRUCache(), v2 != -1: %v", v2)
			return
		}
		cache.Put(4, 4)
		v3 := cache.Get(1)
		if v3 != -1 {
			t.Errorf("Test_LRUCache(), v3 != -1: %v", v3)
			return
		}
		v4 := cache.Get(3)
		if v4 != 3 {
			t.Errorf("Test_LRUCache(), v4 != 3: %v", v4)
			return
		}
		v5 := cache.Get(4)
		if v5 != 4 {
			t.Errorf("Test_LRUCache(), v5 != 4: %v", v4)
			return
		}
	})

	t.Run("case21", func(t *testing.T) {
		cache := ConstructorLRUCache(1)
		cache.Get(6)
		cache.Get(8)
		cache.Put(12, 1)
		cache.Get(2)
		cache.Put(15, 1)
		cache.Put(5, 2)
		cache.Put(1, 15)
		cache.Put(4, 2)
		cache.Get(5)
		cache.Put(15, 15)
	})
}
