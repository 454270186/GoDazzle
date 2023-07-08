package linklist

import (
	"fmt"
	"strings"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list"
)

var _ list.List = (*Linklist)(nil)

type Linklist struct {
	head *element
	tail *element
	size int
}

type element struct {
	value interface{}
	next *element
}

func New(values ...interface{}) *Linklist {
	link := Linklist{}
	link.Add(values...)

	return &link
}

func (l *Linklist) Get(index int) (interface{}, bool) {
	if !l.isInRange(index) {
		return nil, false
	}
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.next
	}

	return cur.value, true
}

func (l *Linklist) Remove(index int) {
	if !l.isInRange(index) {
		return
	}
	
	if index == 0 {
		if l.size == 1 {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
		}
	} else {
		preElement := l.head
		for i := 0; i < index - 1; i++ {
			preElement = preElement.next
		}
		preElement.next = preElement.next.next

		if index == l.size - 1{
			l.tail = preElement
		}
	}

	l.size--
}

func (l *Linklist) Add(values ...interface{}) {
	for _, val := range values {
		newElement := element{value: val}
		if l.size == 0 {
			l.head = &newElement
			l.tail = &newElement
		} else {
			l.tail.next = &newElement
			l.tail = &newElement
		}
		l.size++
	}
}

func (l *Linklist) Contains(value interface{}) bool {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.value == value {
			return true
		}
	}

	return false
}

func (l *Linklist) Sort(cmpFunc cmp.Comparator) {
	if l.size < 2 {
		return
	}

	vals := l.Values()
	cmp.Sort(vals, cmpFunc)

	l.Clear()
	l.Add(vals...)
}

func (l *Linklist) Empty() bool {
	return l.size == 0
}

func (l *Linklist) Size() int {
	return l.size
}

func (l *Linklist) Clear() {
	l.head = nil
	l.tail = nil
	l.size = 0
}

func (l *Linklist) Values() []interface{} {
	cur := l.head
	vals := []interface{}{}
	for cur != nil {
		vals = append(vals, cur.value)
		cur = cur.next
	}

	return vals
}

func (l *Linklist) String() string {
	builder := strings.Builder{}
	builder.WriteString("LinkList: [")
	for cur := l.head; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v", cur.value))
		if cur.next != nil {
			builder.WriteString(" -> ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (l *Linklist) isInRange(index int) bool {
	return index >= 0 && index < l.size
}