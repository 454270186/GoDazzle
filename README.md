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

## Benchmark

Run:
```go
go test -run=NO_TEST -bench . -benchmem  -benchtime 1s ./...
```