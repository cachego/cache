# cache

【[中文文档](./README-zh.md)】

Cache interface definition，provide a simple cache service based on memory，support horizontal expansion

## Feature

*   easy to use

*   support horizontal expansion

*   support key expiration

*   support key tag

## Install

### go module

Use go module directly import：

```go
import "github.com/cachego/cache"
```

### go get

Use go get to install：

```go
go get github.com/cachego/cache
```

## Demo

### 1. cache demo

code reference：

```go
package main

import (
	"fmt"

	"github.com/cachego/cache"
)

func main() {
	key := "key1"
	c := cache.NewInMemoryStrCache()
	c.Set(key, "cache value", 0)
	v, err := c.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v) // cache value
	c.Del(key)
	v, err = c.Get(key)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v) // nil
}
```

output：

````

```text
cache value
<nil>
````

### 2. tag cache demo

code reference：

```go
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
```

output：

```text
[key1 key2]
tag-cache value1
tag-cache value2
[]
<nil>
<nil>

```
