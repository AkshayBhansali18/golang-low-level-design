package evictionpolicy

type EvictionPolicy interface {
	keyAccessed(key string) error
	evictKey() error
}
