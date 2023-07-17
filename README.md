# Go-Dazzle
a Golang datastructure package

## Features
- Multiple data structure supporting
- Some structures are Concurrency Safe
- Embedded JSON Encoder/Decoder

## Getting GoDazzle
Run:
```
go get -u github.com/454270186/GoDazzle
```

### Example
```go
package main

import (
	"fmt"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list/linklist"
)

func main() {
	list := linklist.New() // create a linklist

	// Add value
	list.Add(5, 4, 3, 2, 1)
	fmt.Println(list.Values()...) // [5, 4, 3, 2, 1]

	// Get value by index
	fmt.Println(list.Get(1)) // 4

	// Remove value by index
	list.Remove(3)
	fmt.Println(list.Values()...) // [5, 4, 3, 1]

	// Sort list by built-in comparator
	list.Sort(cmp.IntComparator)
	fmt.Println(list.Values()...) // [1, 3, 4, 5]

	// Sort by custom comparator
	list.Sort(func(a, b interface{}) int {
		num1 := a.(int)
		num2 := b.(int)

		if num1 > num2 {
			return -1
		}

		return 0
	})
	fmt.Println(list.Values()...) // [5, 4, 3, 1]
}
```

## Usage

### Container
All data structures have implement container interface
```go
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []any
	String() string
}
```

### List
- Arraylist
- Linklist
- Doubly linked list

**list interface**
impl Container, JsonCoder, JsonDecoder
```go
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(value interface{}) bool
	Sort(cmpFunc cmp.Comparator)
	
	container.Container
}
```

#### JSON
**Example**
```go
	data := `[1, 2, 3, 4, 5]`
	d := linklist.New()
	d.FromJson([]byte(data)) // equal to json.Unmarshal([]byte(data), d)
	fmt.Println(d.Values()...)
```
Output:
```
1 2 3 4 5
```

<br>

### Map
- GoMap -- native map
- SyncMap -- concurrency safe map

**map interface**
impl Container, JsonCoder, JsonDecoder
```go
type Map interface {
	Put(key, value interface{})
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
	Keys() []interface{}

	container.Container
}
```

<br>

### Queue
- LinkQueue
- LoopQueue
- PriorityQueue
- SyncQueue

**queue interface**
```go
type Queue interface {
	Push(val interface{})
	Pop() bool
	Front() (interface{}, bool)
	Sort(cmpFunc cmp.Comparator)

	container.Container
}
```

<br>

## Benchmark

Run:
```go
go test -run=NO_TEST -bench . -benchmem  -benchtime 1s ./...
```