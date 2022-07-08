package cache_test

import (
	"strconv"
	"testing"
	"time"

	"math/rand"

	"github.com/cachego/cache"
)

func TestNoTagCache(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrTagCache()
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

func TestRandNoTagCache(t *testing.T) {
	key := "key2"
	for i := 0; i < 10; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrTagCache()
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

func BenchmarkRandTagCache(b *testing.B) {
	key := "key2"
	for i := 0; i < b.N; i++ {
		value := strconv.Itoa(rand.Intn(25))
		c := cache.NewInMemoryStrTagCache()
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

func TestTagCache(t *testing.T) {
	key1 := "key1"
	key2 := "key2"
	tag := "tag1"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key1, tag, "value1", 0)
	c.SetWithTag(key2, tag, "value2", 0)
	keys, err := c.GetKeys(tag)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}
}

func TestDelWithTag(t *testing.T) {
	key1 := "key1"
	key2 := "key2"
	tag := "tag1"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key1, tag, "value1", 0)
	c.SetWithTag(key2, tag, "value2", 0)
	// test GetKeys
	keys, err := c.GetKeys(tag)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != 2 {
		t.Errorf("expected 2 keys, got %d", len(keys))
	}
	// test DelWithTag
	err = c.DelWithTag(tag)
	if err != nil {
		t.Error(err)
	}
	keys, err = c.GetKeys(tag)
	if err != nil {
		t.Error(err)
	}
	if len(keys) != 0 {
		t.Errorf("expected 0 keys, got %d", len(keys))
	}
	//
	v1, err := c.Get(key1)
	if err != nil {
		t.Error(err)
	}
	if v1 != nil {
		t.Errorf("expected value to be nil, got '%s'", v1)
	}
	//
	v2, err := c.Get(key2)
	if err != nil {
		t.Error(err)
	}
	if v2 != nil {
		t.Errorf("expected value to be nil, got '%s'", v2)
	}
}

func TestTagIsHit(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key, "tag", "value", 0)
	isHit, err := c.IsHit(key)
	if err != nil {
		t.Error(err)
	}
	if isHit != true {
		t.Error("expected value to be true, got", isHit)
	}
}

func TestTagClear(t *testing.T) {
	key := "key1"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key, "tag", "value", time.Second)
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

func TestClearTag(t *testing.T) {
	key := "key1"
	tag := "tag"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key, tag, "value", time.Second)
	time.Sleep(time.Second * 2)
	err := c.Clear()
	if err != nil {
		t.Error(err)
	}
	keys, err := c.GetKeys(tag)
	if len(keys) != 1 {
		t.Errorf("expected 1 keys, got %d", len(keys))
	}
	err = c.ClearTag()
	if err != nil {
		t.Error(err)
	}
	keys, err = c.GetKeys(tag)
	if len(keys) != 0 {
		t.Errorf("expected 1 keys, got %d", len(keys))
	}
}
