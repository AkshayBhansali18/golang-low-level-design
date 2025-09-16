package storage

import (
	"fmt"
	"golang-low-level-design/multilevelcache/cache"
	"golang-low-level-design/multilevelcache/evictionpolicy"
)

type CacheProvider interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}) error
	Remove(key string) error
}
type cacheProviderImpl struct {
	Cache          cache.Cache
	evictionPolicy evictionpolicy.EvictionPolicy
	nextLevelCache CacheProvider
}

func NewCacheProvider(cache cache.Cache, evictionPolicy evictionpolicy.EvictionPolicy, nextLevelCache CacheProvider) CacheProvider {
	return &cacheProviderImpl{
		Cache:          cache,
		evictionPolicy: evictionPolicy,
		nextLevelCache: nextLevelCache,
	}
}

func (c *cacheProviderImpl) Get(key string) (interface{}, bool) {
	if c.Cache == nil {
		return nil, false
	}
	value, ok := c.Cache.Get(key)
	if ok && value != nil {
		return value, true
	}
	return c.nextLevelCache.Get(key)
}

func (c *cacheProviderImpl) Set(key string, value interface{}) error {
	err := c.Cache.Set(key, value)
	if err != nil {
		return fmt.Errorf("cache provider failed to set key %s: %v", key, err)
	}
	err = c.nextLevelCache.Set(key, value)
	if err != nil {
		return fmt.Errorf("cache provider failed to set key %s: %v", key, err)
	}
	return nil
}

func (c *cacheProviderImpl) Remove(key string) error {
	err := c.Cache.Remove(key)
	if err != nil {
		return fmt.Errorf("cache provider failed to remove key %s: %v", key, err)
	}
	err = c.nextLevelCache.Remove(key)
	if err != nil {
		return fmt.Errorf("cache provider failed to remove key %s: %v", key, err)
	}
	return nil
}
