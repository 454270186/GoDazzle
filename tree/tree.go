package tree

import "github.com/454270186/GoDazzle/container"


/* 	Container interface:
Empty() bool
Size() int
Clear()
Values() []any
String() string
*/
type Tree interface {
	container.Container
}