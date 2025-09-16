package evictionpolicy

type LRUEvictionPolicy struct {
}

func (p *LRUEvictionPolicy) keyAccessed(key string) error {
	// implement this
	return nil
}

func (p *LRUEvictionPolicy) evictKey() error {
	// implement this
	return nil
}
