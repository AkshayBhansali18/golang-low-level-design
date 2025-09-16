package storage

type nullEffectLevelCache struct {
}

func NewNullEffectLevelCache() CacheProvider {
	return &nullEffectLevelCache{}
}

func (c *nullEffectLevelCache) Get(key string) (interface{}, bool) {
	return nil, false
}

func (c *nullEffectLevelCache) Set(key string, value interface{}) error {
	return nil
}

func (c *nullEffectLevelCache) Remove(key string) error {
	return nil
}
