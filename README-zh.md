# cache

cache 接口定义，提供一个基于内存的简单服务，支持横向扩展。

## 特点

- 简单易用
- 支持横向扩展
- 支持 key 过期
- 支持 key 打 tag

## 安装

### go module
使用 go module 直接 import ：

```go
import "github.com/cachego/cache"
```

### go get
使用 go get 安装：

```go
go get github.com/cachego/cache
```

## Demo

[完整例子](https://github.com/cachego/cache/tree/main/example)
### 1. cache demo

参考代码：

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

输出：

```text
cache value
<nil>
```


### 2. tag cache demo

参考代码：

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
输出：

```text
[key1 key2]
tag-cache value1
tag-cache value2
[]
<nil>
<nil>
```
