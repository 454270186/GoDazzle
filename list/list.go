package list

import "github.com/454270186/GoDazzle/container"

// Common interface of all types of list
type List interface {
	Get(index int) (interface{}, bool)
	Remove(index int)
	Add(values ...interface{})
	Contains(values ...interface{}) bool
	Sort()
	
	container.Container
}