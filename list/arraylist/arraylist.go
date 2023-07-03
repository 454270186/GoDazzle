package arraylist

import "fmt"

type ArrayList struct {
	elements []interface{}
	size int
}

var extendFactor = 2

func New(values ...any) *ArrayList {
	arrayList := &ArrayList{}

	if len(values) > 0 {
		arrayList.Add(values...)
	}

	return arrayList
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
}

func (arrayList *ArrayList) Add(values ...interface{}) {
	arrayList.extend(len(values))

	for _, val := range values {
		arrayList.elements[arrayList.size] = val
		arrayList.size++
	}
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
	return index >= 0 && arrayList.size > index
}

func (arrList *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap)
	copy(newElements, arrList.elements)
	arrList.elements = newElements
}

func (arrayList *ArrayList) extend(n int) {
	if cap(arrayList.elements) == 0 {
		arrayList.resize(n)
	}

	curCapacity := cap(arrayList.elements)
	if len(arrayList.elements) + n > curCapacity {
		newCap := curCapacity * extendFactor
		arrayList.resize(newCap)
	}
}