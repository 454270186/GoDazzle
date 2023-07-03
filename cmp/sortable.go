package cmp

import (
	"sort"
)

// Sort is the Open interface implementing go "sort" interface
func Sort(values []interface{}, cmp Comparator) {
	sort.Sort(sortable{values, cmp})
}

type sortable struct {
	values []interface{}
	cmp Comparator
}

func (s sortable) Len() int {
	return len(s.values)
}

func (s sortable) Swap(i, j int) {
	s.values[i], s.values[j] = s.values[j], s.values[i]
}

func (s sortable) Less(i, j int) bool {
	return s.cmp(s.values[i], s.values[j]) < 0
}