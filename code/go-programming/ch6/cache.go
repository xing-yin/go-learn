package ch6

import "sync"

// 实现方式1：简单的 cache
var (
	mu      sync.Mutex
	mapping = make(map[string]string)
)

func lookup(key string) string {
	mu.Lock()
	defer mu.Unlock()
	v := mapping[key]
	return v
}

// 更优的实现: 将两个包级的变量放在了 cache 这个struct一组内
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookupV2(key string) string {
	cache.Lock()
	defer cache.Unlock()
	v := cache.mapping[key]
	return v
}
