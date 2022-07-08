package cache_test

import (
	"strconv"
	"testing"
	"time"

	"math/rand"

	"github.com/cachego/cache"
)

func TestCache(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrCache()
	c.Set(key, "value", 0)
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
	if v != nil {
		t.Errorf("expected value to be nil, got '%s'", v)
	}
}

func TestRandCache(t *testing.T) {
	key := "key2"
	for i := 0; i < 10; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrCache()
		c.Set(key, value, 0)
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
		if v != nil {
			t.Errorf("expected value to be nil, got '%s'", v)
		}
	}
}

func BenchmarkRandCache(b *testing.B) {
	key := "key2"
	for i := 0; i < b.N; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrCache()
		c.Set(key, value, 0)
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
		if v != nil {
			b.Errorf("expected value to be nil, got '%s'", v)
		}
	}
}

func TestIsHit(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrCache()
	c.Set(key, "value", 0)
	isHit, err := c.IsHit(key)
	if err != nil {
		t.Error(err)
	}
	if isHit != true {
		t.Error("expected value to be true, got", isHit)
	}
}

func TestClear(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrCache()
	c.Set(key, "value", time.Second)
	time.Sleep(time.Second * 2)
	err := c.Clear()
	if err != nil {
		t.Error(err)
	}
	isHit, err := c.IsHit(key)
	if err != nil {
		t.Error(err)
	}
	if isHit != false {
		t.Error("expected value to be false, got", isHit)
	}
}
