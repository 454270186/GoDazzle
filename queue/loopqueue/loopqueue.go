package loopqueue

import (
	"math"

	"github.com/454270186/GoDazzle/cmp"
	"github.com/454270186/GoDazzle/queue"
)

var _ queue.Queue = (*LoopQueue)(nil)

type LoopQueue struct {
	q     []interface{}
	begin int
	end   int
	cap   int
}

// cap is the capacity of queue
// default cap is 10
func New(cap ...int) *LoopQueue {
	if len(cap) > 0 {
		return &LoopQueue{
			q: make([]interface{}, cap[0]),
			begin: 0,
			end: 0,
			cap: cap[0],
		}
	}

	return &LoopQueue{
		q: make([]interface{}, 10),
		begin: 0,
		end: 0,
		cap: 10,
	}
}

// LoopQueue is cap-limited,
// When queue is full, call Push() will panic
func (lq *LoopQueue) Push(val interface{}) {
	if (lq.end + 1) % lq.cap == lq.begin {
		panic("queue is full")
	}

	lq.q[lq.end] = val
	lq.end = (lq.end + 1) % lq.cap
}

func (lq *LoopQueue) Pop() bool {
	if lq.begin == lq.end {
		return false
	}

	lq.begin = (lq.begin + 1) % lq.cap
	return true
}

func (lq *LoopQueue) Front() (interface{}, bool) {
	if lq.begin == lq.end {
		return nil, false
	}

	val := lq.q[lq.begin]
	return val, true
}

func (lq *LoopQueue) Sort(cmpFunc cmp.Comparator) {
	cmp.Sort(lq.q, cmpFunc)
}

func (lq *LoopQueue) Empty() bool {
	return lq.begin == lq.end
}

func (lq *LoopQueue) Size() int {
	return int(math.Abs(float64(lq.end - lq.begin)))
}

func (lq *LoopQueue) Clear() {
	lq.begin = 0
	lq.end = 0
}

func (lq *LoopQueue) Values() []any {
	vals := make([]any, 0)
	i := lq.begin
	for i != lq.end {
		vals = append(vals, lq.q[i])
		i = (i + 1) % lq.cap
	}

	return vals
}

func (lq *LoopQueue) String() string {
	return "not impl"
}