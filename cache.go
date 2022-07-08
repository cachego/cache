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

	// Check the key is hit, if not, return false, otherwise return true.
	IsHit(key string) (bool, error)

	// Clear the expiration value
	Clear() error
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
		if v.Exp != 0 && v.Exp < time.Now().Unix() {
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
	if ttl == 0 {
		c.data[key] = inMemoryStrCacheItem{
			Val: val.(string),
			Exp: 0,
		}
	} else {
		c.data[key] = inMemoryStrCacheItem{
			Val: val.(string),
			Exp: time.Now().Unix() + int64(ttl.Seconds()),
		}
	}
	return
}

func (c *inMemoryStrCache) IsHit(key string) (isHit bool, err error) {
	val, err := c.Get(key)
	if err != nil {
		return
	}
	return val != nil, nil
}

func (c *inMemoryStrCache) Clear() (err error) {
	for key, _ := range c.data {
		if _, err = c.Get(key); err != nil {
			return
		}
	}
	return
}
