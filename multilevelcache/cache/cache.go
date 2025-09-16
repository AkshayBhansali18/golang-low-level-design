package cache

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string) (interface{}, bool)
	Remove(key string) error
}
