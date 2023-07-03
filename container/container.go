package container

// the core functions of all data structures
type Container interface {
	Empty() bool
	Size() int
	Clear()
	Values() []any
	String() string
}
