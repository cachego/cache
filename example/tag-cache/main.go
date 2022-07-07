package main

import (
	"fmt"

	"github.com/cachego/cache"
)

func main() {
	key1 := "key1"
	key2 := "key2"
	tag := "tag1"
	c := cache.NewInMemoryStrTagCache()
	c.SetWithTag(key1, tag, "tag-cache value1", 0)
	c.SetWithTag(key2, tag, "tag-cache value2", 0)
	keys, _ := c.GetKeys(tag)
	fmt.Println(keys) // [key1 key2]
	v1, _ := c.Get(key1)
	fmt.Println(v1) // tag-cache value1
	v2, _ := c.Get(key2)
	fmt.Println(v2) // tag-cache value2

	c.DelWithTag(tag)
	keys, _ = c.GetKeys(tag)
	fmt.Println(keys) // []
	v11, _ := c.Get(key1)
	fmt.Println(v11) // nil
	v22, _ := c.Get(key2)
	fmt.Println(v22) // nil
}
