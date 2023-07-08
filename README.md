# Go-Dazzle
a Golang datastructure package

## Features
- Multiple data structure supporting
- Concurrency Safe
- Embedded JSON Encoder/Decoder

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
impl Container interface
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
