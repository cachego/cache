package cache

// cache is an interface for caching data.
type Cache interface {
	// Get returns the value for the given key. If the key does not exist, returns nil.
	Get(key string) (interface{}, error)

	//delete the value for the given key. if the key does not exist, returns EmptyKeyError.
	Del(key string) error

	// Set sets the value for the given key.
	Set(key string, v interface{}) error
}

// NewInMemoryStrCache returns a new in-memory cache.
func NewInMemoryStrCache() Cache {
	return &inMemoryStrCache{
		data: make(map[string]string),
	}
}

type inMemoryStrCache struct {
	data map[string]string
}

func (c *inMemoryStrCache) Get(key string) (v interface{}, err error) {
	v = c.data[key]
	return
}

func (c *inMemoryStrCache) Del(key string) (err error) {
	delete(c.data, key)
	return nil
}

func (c *inMemoryStrCache) Set(key string, v interface{}) (err error) {
	c.data[key] = v.(string)
	return nil
}
