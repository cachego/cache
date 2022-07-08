package cache

import "time"

// TagCache is an interface for caching data.
type TagCache interface {
	Cache

	// Set sets the value for the given key and tag.
	SetWithTag(key string, tag string, v interface{}, ttl time.Duration) error

	// GetKeys returns the keys for the given tag.
	GetKeys(tag string) ([]string, error)

	// DelWithTag the value for the given tag.
	DelWithTag(tag string) error

	// clear not include a valid key tag
	ClearTag() error
}

// NewInMemoryStrCache returns a new in-memory cache.
func NewInMemoryStrTagCache() TagCache {
	return &inMemoryStrTagCache{
		data: make(map[string]inMemoryStrTagCacheItem),
		tag:  make(map[string]tagValue),
	}
}

type inMemoryStrTagCache struct {
	data map[string]inMemoryStrTagCacheItem
	tag  map[string]tagValue
}

type inMemoryStrTagCacheItem struct {
	Val string
	Exp int64
}

type tagValue struct {
	keyMap map[string]bool
}

func (c *inMemoryStrTagCache) GetKeys(tag string) (keys []string, err error) {
	for k := range c.tag[tag].keyMap {
		keys = append(keys, k)
	}
	return
}

func (c *inMemoryStrTagCache) SetWithTag(key string, tag string, v interface{}, ttl time.Duration) (err error) {
	c.Set(key, v, ttl)
	var tv tagValue
	if c.tag[tag].keyMap != nil {
		tv = c.tag[tag]
	} else {
		tv = tagValue{keyMap: make(map[string]bool)}
	}
	tv.keyMap[key] = true
	c.tag[tag] = tv
	return
}

func (c *inMemoryStrTagCache) Get(key string) (val interface{}, err error) {
	if v, ok := c.data[key]; ok {
		if v.Exp != 0 && v.Exp < time.Now().Unix() {
			val = nil
		} else {
			val = v.Val
		}
	} else {
		val = nil
	}
	return
}

func (c *inMemoryStrTagCache) Del(key string) (err error) {
	delete(c.data, key)
	return
}

func (c *inMemoryStrTagCache) DelWithTag(tag string) (err error) {
	keys, err := c.GetKeys(tag)
	delete(c.tag, tag)
	if err == nil {
		for _, key := range keys {
			c.Del(key)
		}
	}
	return
}

func (c *inMemoryStrTagCache) Set(key string, val interface{}, ttl time.Duration) (err error) {
	if ttl == 0 {
		c.data[key] = inMemoryStrTagCacheItem{
			Val: val.(string),
			Exp: 0,
		}
	} else {
		c.data[key] = inMemoryStrTagCacheItem{
			Val: val.(string),
			Exp: time.Now().Unix() + int64(ttl.Seconds()),
		}
	}
	return
}

func (c *inMemoryStrTagCache) IsHit(key string) (isHit bool, err error) {
	val, err := c.Get(key)
	if err != nil {
		return
	}
	return val != nil, nil
}

func (c *inMemoryStrTagCache) Clear() (err error) {
	for key, _ := range c.data {
		if _, err = c.Get(key); err != nil {
			return
		}
	}
	return
}

func (c *inMemoryStrTagCache) ClearTag() error {
	for tag, tagValue := range c.tag {
		isValid := false
		for key, _ := range tagValue.keyMap {
			isHit, err := c.IsHit(key)
			if err != nil {
				return err
			}
			isValid = isValid || isHit
		}
		if !isValid {
			c.DelWithTag(tag)
		}
	}
	return nil
}
