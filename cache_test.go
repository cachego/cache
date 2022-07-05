package cache_test

import (
	"strconv"
	"testing"

	"math/rand"

	"github.com/cachego/cache"
)

func TestCache(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrCache()
	c.Set(key, "value")
	v, err := c.Get(key)
	if err != nil {
		t.Error(err)
	}
	if v != "value" {
		t.Errorf("expected value to be 'value', got '%s'", v)
	}
	c.Del(key)
	v, err = c.Get(key)
	if err != nil {
		t.Error(err)
	}
	if v == nil {
		t.Errorf("expected value to be nil, got '%s'", v)
	}
}

func TestRandCache(t *testing.T) {
	key := "key2"
	for i := 0; i < 10; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrCache()
		c.Set(key, value)
		v, err := c.Get(key)
		if err != nil {
			t.Error(err)
		}
		if v != value {
			t.Errorf("expected value to be 'value', got '%s'", v)
		}
		c.Del(key)
		v, err = c.Get(key)
		if err != nil {
			t.Error(err)
		}
		if v == nil {
			t.Errorf("expected value to be nil, got '%s'", v)
		}
	}
}

func BenchmarkRandCache(b *testing.B) {
	key := "key2"
	for i := 0; i < b.N; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrCache()
		c.Set(key, value)
		v, err := c.Get(key)
		if err != nil {
			b.Error(err)
		}
		if v != value {
			b.Errorf("expected value to be 'value', got '%s'", v)
		}
		c.Del(key)
		v, err = c.Get(key)
		if err != nil {
			b.Error(err)
		}
		if v == nil {
			b.Errorf("expected value to be nil, got '%s'", v)
		}
	}
}
