package stack

import "github.com/454270186/GoDazzle/container"

/* 	Container interface:
Empty() bool
Size() int
Clear()
Values() []any
String() string
*/

type Stack interface {
	Push(val interface{})
	Pop()
	Top() (interface{}, bool)

	container.Container
}