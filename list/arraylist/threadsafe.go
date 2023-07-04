package arraylist

import (
	"fmt"
	"strings"
	"sync"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list"
)

var _ list.List = (*TFArrList)(nil)

type TFArrList struct {
	sync.RWMutex
	elements []interface{}
}

func newTFArrList() *TFArrList {
	return &TFArrList{}
}

func (tf *TFArrList) Get(index int) (interface{}, bool) {
	tf.RLock()
	defer tf.RUnlock()

	if !tf.isInRange(index) {
		return nil, false
	}

	return tf.elements[index], true
}

func (tf *TFArrList) Remove(index int) {
	tf.Lock()
	if !tf.isInRange(index) {
		return
	}

	tf.elements[index] = nil
	copy(tf.elements[index:], tf.elements[index+1:])

	tf.Unlock()

	//tf.shrink()
}

func (tf *TFArrList) Add(values ...interface{}) {
	tf.Lock()
	defer tf.Unlock()
	//tf.extend(len(values))

	tf.elements = append(tf.elements, values...)
}

func (tf *TFArrList) Contains(value interface{}) bool {
	tf.RLock()
	defer tf.RUnlock()

	isContains := false

	for _, val := range tf.elements {
		if val == value {
			isContains = true
		}
	}

	return isContains
}

func (tf *TFArrList) Sort(cmpFunc cmp.Comparator) {
	tf.Lock()
	defer tf.Unlock()
	cmp.Sort(tf.elements, cmpFunc)
}

func (tf *TFArrList) Empty() bool {
	tf.RLock()
	defer tf.RUnlock()
	return len(tf.elements) == 0
}

func (tf *TFArrList) Size() int {
	tf.RLock()
	defer tf.RUnlock()
	return len(tf.elements)
}

func (tf *TFArrList) Clear() {
	tf.Lock()
	defer tf.Unlock()
	tf.elements = make([]interface{}, 0)
}

func (tf *TFArrList) Values() []interface{} {
	tf.Lock()
	defer tf.Unlock()

	vals := make([]interface{}, len(tf.elements))
	copy(vals, tf.elements)
	
	return vals
}

func (tf *TFArrList) String() string {
	builder := strings.Builder{}
	builder.WriteString("ArrayList: ")
	builder.WriteString("[")

	tf.RLock()
	for i, val := range tf.elements {
		if val == nil {
			continue
		}
		valStr := fmt.Sprintf("%v", val)
		builder.WriteString(valStr)
		if i != len(tf.elements)-1 && tf.elements[i+1] != nil{
			builder.WriteString(", ")
		}
	}
	tf.RUnlock()

	builder.WriteString("]")
	return builder.String()
}

func (tf *TFArrList) Println() {
	for _, val := range tf.elements {
		if val == nil {
			continue
		}
		fmt.Print(val)
		fmt.Print(" ")
	}

	fmt.Println()
}

// No Lock inside
func (tf *TFArrList) isInRange(index int) bool {
	return index >= 0 && len(tf.elements) > index
}

// // No Lock inside
// func (tf *TFArrList) resize(cap int) {
// 	newElements := make([]interface{}, cap)
// 	copy(newElements, tf.elements)
// 	tf.elements = newElements
// }

// // No Lock inside
// func (tf *TFArrList) extend(n int) {
// 	if len(tf.elements) == 0 {
// 		tf.resize(n)
// 	}

// 	curCapacity := cap(tf.elements)
// 	if len(tf.elements) + n > curCapacity {
// 		newCap := curCapacity * extendFactor
// 		tf.resize(newCap)
// 	}
// }

// func (tf *TFArrList) shrink() {
// 	if shrinkFactor == 0 {
// 		return
// 	}
// 	tf.Lock()
// 	defer tf.Unlock()
// 	curCapacity := cap(tf.elements)
// 	if len(tf.elements) < curCapacity * int(shrinkFactor) {
// 		tf.resize(len(tf.elements))
// 	}
// }

