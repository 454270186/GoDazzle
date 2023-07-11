package arraylist

import (
	"fmt"
	"strings"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list"
)

var _ list.List = (*ArrayList)(nil)

type ArrayList struct {
	elements []interface{}
	size int
}

var (
	shrinkFactor = 0.25
)

func New(values ...interface{}) *ArrayList {
	a := ArrayList{}
	a.Add(values...)
	return &a
}

func (arrayList *ArrayList) Get(index int) (interface{}, bool) {
	if !arrayList.isInRange(index) {
		return nil, false
	}

	return arrayList.elements[index], true
}

func (arrayList *ArrayList) Remove(index int) {
	if !arrayList.isInRange(index) {
		return
	}

	arrayList.elements[index] = nil
	copy(arrayList.elements[index:], arrayList.elements[index+1:])
	arrayList.size--

	arrayList.shrink()
}

func (arrayList *ArrayList) Add(values ...interface{}) {
	for _, val := range values {
		arrayList.elements = append(arrayList.elements, val)
		arrayList.size++
	}
}

func (arrayList *ArrayList) Contains(value interface{}) bool {
	isContains := false

	for _, val := range arrayList.elements {
		if val == value {
			isContains = true
		}
	}

	return isContains
}

func (arrayList *ArrayList) Sort(cmpFunc cmp.Comparator) {
	cmp.Sort(arrayList.elements[:arrayList.size], cmpFunc)
}

func (arrayList *ArrayList) Empty() bool {
	return arrayList.size == 0
}

func (arrayList *ArrayList) Size() int {
	return arrayList.size
}

func (arrayList *ArrayList) Clear() {
	arrayList.elements = make([]interface{}, 0)
	arrayList.size = 0
}

func (arrayList *ArrayList) Values() []interface{} {
	return append([]interface{}{}, arrayList.elements...)
}

func (arrayList *ArrayList) String() string {
	builder := strings.Builder{}
	builder.WriteString("ArrayList: ")
	builder.WriteString("[")
	for i, val := range arrayList.elements {
		if val == nil {
			continue
		}
		valStr := fmt.Sprintf("%v", val)
		builder.WriteString(valStr)
		if i != len(arrayList.elements)-1 && arrayList.elements[i+1] != nil{
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (arrayList *ArrayList) Println() {
	for _, val := range arrayList.elements {
		if val == nil {
			continue
		}
		fmt.Print(val)
		fmt.Print(" ")
	}

	fmt.Println()
}

func (arrayList *ArrayList) isInRange(index int) bool {
	return index >= 0 && index < arrayList.size
}

func (arrList *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap)
	copy(newElements, arrList.elements)
	arrList.elements = newElements
}

func (arrayList *ArrayList) shrink() {
	if shrinkFactor == 0 {
		return
	}

	curCapacity := cap(arrayList.elements)
	if arrayList.size < curCapacity * curCapacity {
		arrayList.resize(arrayList.size)
	}
}