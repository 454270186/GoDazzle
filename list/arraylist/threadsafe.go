package arraylist

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list"
)

var _ list.List = (*TFArrList)(nil)

type TFArrList struct {
	sync.Mutex
	elements []interface{}
	size atomic.Int32
}

func newTFArrList() *TFArrList {
	return &TFArrList{}
}

func (tf *TFArrList) Get(index int) (interface{}, bool) {
	if !tf.isInRange(index) {
		return nil, false
	}

	tf.Lock()
	defer tf.Unlock()

	return tf.elements[index], true
}

func (tf *TFArrList) Remove(index int) {
	if !tf.isInRange(index) {
		return
	}

	tf.Lock()
	{
		tf.elements[index] = nil
		copy(tf.elements[index:], tf.elements[index+1:])
		tf.size.Add(-1)
	}
	tf.Unlock()

	tf.shrink()
}

func (tf *TFArrList) Add(values ...interface{}) {
	tf.extend(len(values))

	tf.Lock()
	defer tf.Unlock()
	for _, val := range values {
		tf.elements[tf.size.Load()] = val
		tf.size.Add(1)
	}
}

func (tf *TFArrList) Contains(value interface{}) bool {
	tf.Lock()
	defer tf.Unlock()

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
	return tf.size.Load() == 0
}

func (tf *TFArrList) Size() int {
	return int(tf.size.Load())
}

func (tf *TFArrList) Clear() {
	tf.Lock()
	defer tf.Unlock()
	tf.elements = make([]interface{}, 0)
	tf.size.Store(0)
}

func (tf *TFArrList) Values() []interface{} {
	tf.Lock()
	defer tf.Unlock()

	vals := make([]interface{}, 0, tf.size.Load())
	nilIndex := -1
	for i, val := range tf.elements {
		if val == nil {
			nilIndex = i
			break
		}
	}
	if nilIndex >= 0 {
		copy(vals, tf.elements[0:nilIndex])
	}
	
	return vals
}

func (tf *TFArrList) String() string {
	builder := strings.Builder{}
	builder.WriteString("ArrayList: ")
	builder.WriteString("[")

	tf.Lock()
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
	tf.Unlock()

	builder.WriteString("]\n")
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

func (tf *TFArrList) isInRange(index int) bool {
	return index >= 0 && int(tf.size.Load()) > index
}

func (tf *TFArrList) resize(cap int) {
	tf.Lock()
	defer tf.Unlock()

	newElements := make([]interface{}, cap)
	copy(newElements, tf.elements)
	tf.elements = newElements
}

func (tf *TFArrList) extend(n int) {
	tf.Lock()
	curCapacity := cap(tf.elements)
	tf.Unlock()
	if curCapacity == 0 {
		tf.resize(n)
		curCapacity = n
	}

	if int(tf.size.Load()) + n > curCapacity {
		newCap := curCapacity * extendFactor
		tf.resize(newCap)
	}
}

func (tf *TFArrList) shrink() {
	if shrinkFactor == 0 {
		return
	}
	tf.Lock()
	curCapacity := cap(tf.elements)
	tf.Unlock()
	if int(tf.size.Load()) < curCapacity * int(shrinkFactor) {
		tf.resize(int(tf.size.Load()))
	}
}

