package list

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

// Common interface of all types of list
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(value interface{}) bool
	Sort(cmpFunc cmp.Comparator)
	
	container.Container
}