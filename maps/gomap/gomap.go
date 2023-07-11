package gomap

import (
	"fmt"

	"github.com/454270186/GoDazzle/maps"
)

/*
	Encapsulation of Go native map
*/

var _ maps.Map = (*GoMap)(nil)

type GoMap struct {
	m map[any]any
}

func New() *GoMap {
	return &GoMap{
		m: make(map[any]any),
	}
}

func (g *GoMap) Put(key, value interface{}) {
	g.m[key] = value
}

func (g *GoMap) Get(key interface{}) (interface{}, bool) {
	val, ok := g.m[key]
	return val, ok
}

func (g *GoMap) Remove(key interface{}) {
	delete(g.m, key)
}

func (g *GoMap) Keys() []interface{} {
	keys := make([]any, g.Size())
	i := 0
	for k := range g.m {
		keys[i] = k
		i++
	}

	return keys
}

func (g *GoMap) Empty() bool {
	return g.Size() == 0
}

func (g *GoMap) Size() int {
	return len(g.m)
}

func (g *GoMap) Clear() {
	g.m = make(map[any]any)
}

func (g *GoMap) Values() []any {
	vals := make([]interface{}, g.Size())
	i := 0
	for _, v := range g.m {
		vals[i] = v
		i++
	}

	return vals
}

func (g *GoMap) String() string {
	str := "GoMap: "
	str += fmt.Sprintf("%v", g.m)
	return str
}
