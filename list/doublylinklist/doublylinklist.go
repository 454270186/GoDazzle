package doublylinklist

import (
	"fmt"
	"strings"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list"
)

var _ list.List = (*DoublyLinkList)(nil)

type DoublyLinkList struct {
	head *element
	tail *element
	size int
}

type element struct {
	value interface{}
	pre *element
	next *element
}

func New(values ...interface{}) *DoublyLinkList {
	dbLinklist := DoublyLinkList{}
	dbLinklist.Add(values...)
	return &dbLinklist
}

func (d *DoublyLinkList) Get(index int) (interface{}, bool) {
	if !d.isInRange(index) {
		return nil, false
	}

	if index <= d.size / 2 {
		cur := d.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		return cur.value, true

	} else {
		cur := d.tail
		for i := (d.size-1) - index; i > 0; i-- {
			cur = cur.pre
		}
		return cur.value, true
	}
}

func (d *DoublyLinkList) Remove(index int) {
	if !d.isInRange(index) {
		return
	}

	if index == 0 {
		if d.size == 1 {
			d.head = nil
			d.tail = nil
		} else {
			d.head = d.head.next
		}
	} else if index <= d.size / 2 {
		cur := d.head
		for i := 0; i < index; i++ {
			cur = cur.next
		}
		cur.pre.next = cur.next
		cur.next.pre = cur.pre
	} else if index > d.size / 2 {
		cur := d.tail
		for i := (d.size-1) - index; i > 0; i-- {
			cur = cur.pre
		}
		cur.pre.next = cur.next
		cur.next.pre = cur.pre
	}

	d.size--
}

func (d *DoublyLinkList) Add(values ...interface{}) {
	for _, val := range values {
		newElement := element{value: val}
		if d.size == 0 {
			d.head = &newElement
			d.tail = &newElement
		} else {
			newElement.pre = d.tail
			d.tail.next = &newElement
			d.tail = &newElement
		}
		d.size++
	}
}

func (d *DoublyLinkList) Contains(value interface{}) bool {
	l, r := d.head, d.tail
	for l != r && r.next != l && l.pre != r {
		if l.value == value || r.value == value {
			return true
		}
		l = l.next
		r = r.pre
	}

	return false
}

func (d *DoublyLinkList) Sort(cmpFunc cmp.Comparator) {
	if d.size < 2 {
		return
	}

	vals := d.Values()
	cmp.Sort(vals, cmpFunc)
	d.Clear()
	d.Add(vals...)
}

func (d *DoublyLinkList) Empty() bool {
	return d.size == 0
}

func (d *DoublyLinkList) Size() int {
	return d.size
}

func (d *DoublyLinkList) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *DoublyLinkList) Values() []interface{} {
	cur := d.head
	vals := []interface{}{}
	for cur != nil {
		vals = append(vals, cur.value)
		cur = cur.next
	}

	return vals
}

func (d *DoublyLinkList) String() string {
	builder := strings.Builder{}
	builder.WriteString("DoublyLinkList: [")
	for cur := d.head; cur != nil; cur = cur.next {
		builder.WriteString(fmt.Sprintf("%v", cur.value))
		if cur.next != nil {
			builder.WriteString(" -> ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (d *DoublyLinkList) isInRange(index int) bool {
	return index >= 0 && index < d.size
}