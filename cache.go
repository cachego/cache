package cache

import "time"

// cache is an interface for caching data.
type Cache interface {
	// Get returns the value for the given key.
	Get(key string) (interface{}, error)

	//delete the value for the given key.
	Del(key string) error

	// Set sets the value for the given key. If ttl is 0, the value will not expire
	Set(key string, val interface{}, ttl time.Duration) error
}

// NewInMemoryStrCache returns a new in-memory cache.
func NewInMemoryStrCache() Cache {
	return &inMemoryStrCache{
		data: make(map[string]inMemoryStrCacheItem),
	}
}

type inMemoryStrCache struct {
	data map[string]inMemoryStrCacheItem
}

type inMemoryStrCacheItem struct {
	Val string
	Exp int64
}

func (c *inMemoryStrCache) Get(key string) (val interface{}, err error) {
	if v, ok := c.data[key]; ok {
		if v.Exp > time.Now().Unix() {
			c.Del(key)
			val = nil
		} else {
			val = v.Val
		}
	} else {
		val = nil
	}
	return
}

func (c *inMemoryStrCache) Del(key string) (err error) {
	delete(c.data, key)
	return
}

func (c *inMemoryStrCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	c.data[key] = inMemoryStrCacheItem{
		Val: val.(string),
		Exp: time.Now().Unix() + int64(ttl.Seconds()),
	}
	return
}
