package cache

type inMemoryCache struct {
	cache map[string]interface{}
}

func NewInMemoryCache() Cache {
	cache := make(map[string]interface{})
	return &inMemoryCache{cache: cache}
}

func (c *inMemoryCache) Get(key string) (interface{}, bool) {
	if _, ok := c.cache[key]; ok {
		return c.cache[key], true
	}
	return nil, false
}

func (c *inMemoryCache) Set(key string, value interface{}) error {
	c.cache[key] = value
	return nil
}

func (c *inMemoryCache) Remove(key string) error {
	delete(c.cache, key)
	return nil
}
