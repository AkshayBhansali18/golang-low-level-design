package main

import (
	"fmt"
	"golang-low-level-design/multilevelcache/cache"
	"golang-low-level-design/multilevelcache/evictionpolicy"
	"golang-low-level-design/multilevelcache/storage"
)

func main() {
	l1Cache := cache.NewInMemoryCache()
	l1Cache.Set("key2", "value2")
	l1CacheEvictionPolicy := &evictionpolicy.LRUEvictionPolicy{}
	l2Cache := cache.NewInMemoryCache()
	l2Cache.Set("key1", "value1")
	l2CacheEvictionPolicy := &evictionpolicy.LRUEvictionPolicy{}
	nilEffectCache := storage.NewNullEffectLevelCache()
	cacheProviderL2 := storage.NewCacheProvider(l2Cache, l2CacheEvictionPolicy, nilEffectCache)
	cacheProviderL1 := storage.NewCacheProvider(l1Cache, l1CacheEvictionPolicy, cacheProviderL2)
	val, ok := cacheProviderL1.Get("key1")
	// value is expected in multilevel cache
	fmt.Println(val, ok)
	val, ok = cacheProviderL1.Get("key2")
	// value is expected in multilevel cache
	fmt.Println(val, ok)
	val, ok = cacheProviderL1.Get("key3")
	// value is not expected in multilevel cache
	fmt.Println(val, ok)
}
