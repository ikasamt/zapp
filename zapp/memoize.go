package zapp

import (
	"sync"
	"time"
)

var mu sync.RWMutex = sync.RWMutex{}
var cache map[string]interface{} = make(map[string]interface{})

// Set
func MemoizeSet(key string, val interface{}, expire time.Duration) {
	mu.Lock()
	defer mu.Unlock()

	cache[key] = val

	time.AfterFunc(expire, func() {
		mu.Lock()
		defer mu.Unlock()
		delete(cache, key)
	})
}

// Get
func MemoizeGet(key string) (value interface{}, ok bool) {
	mu.RLock()
	defer mu.RUnlock()

	value, ok = cache[key]
	return
}
