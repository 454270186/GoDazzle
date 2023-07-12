package queue

import (
	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/container"
)

/* 	Container interface:
Empty() bool
Size() int
Clear()
Values() []any
String() string
*/

type Queue interface {
	Push(val interface{})
	Pop() bool
	Front() (interface{}, bool)
	Sort(cmpFunc cmp.Comparator)

	container.Container
}