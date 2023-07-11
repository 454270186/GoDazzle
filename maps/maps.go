package maps

import "github.com/454270186/GoDazzle/container"

/* 	Container interface:
Empty() bool
Size() int
Clear()
Values() []any
String() string
*/

type Map interface {
	Put(key, value interface{})
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{})
	Keys() []interface{}
	container.Container
}