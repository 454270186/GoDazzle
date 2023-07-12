package linkqueue

import (
	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/list/linklist"
	"github.com/454270186/GoDazzle/queue"
)

var _ queue.Queue = (*LinkQueue)(nil)

type LinkQueue struct {
	items *linklist.Linklist
}

func New(values ...interface{}) *LinkQueue {
	return &LinkQueue{
		items: linklist.New(values...),
	}
}

func (q *LinkQueue) Push(val interface{}) {
	q.items.Add(val)
}

func (q *LinkQueue) Pop() bool {
	if q.Empty() {
		return false
	}

	q.items.Remove(0)
	return true
}

func (q *LinkQueue) Front() (interface{}, bool) {
	val, ok := q.items.Get(0)
	if !ok {
		return nil, false
	}

	return val, true
}

func (q *LinkQueue) Sort(cmpFunc cmp.Comparator) {
	q.items.Sort(cmpFunc)
}

func (q *LinkQueue) Empty() bool {
	return q.items.Empty()
}

func (q *LinkQueue) Size() int {
	return q.items.Size()
}

func (q *LinkQueue) Clear() {
	q.items.Clear()
}

func (q *LinkQueue) Values() []any {
	return q.items.Values()
}

func (q *LinkQueue) String() string {
	return "not impl"
}