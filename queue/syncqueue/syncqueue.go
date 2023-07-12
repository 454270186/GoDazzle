package syncqueue

import (
	"time"
)

type SyncQueue struct {
	q chan interface{}
	cap int
	size int
}

func New(cap ...int) *SyncQueue {
	if len(cap) > 0 {
		return &SyncQueue{
			q: make(chan interface{}, cap[0]),
			cap: cap[0],
			size: 0,
		}
	}

	return &SyncQueue{
		q: make(chan interface{}, 10),
		cap: 10,
		size: 0,
	}
}

// When queue is full, call Push() will block goroutine
// Until queue is not full
func (s *SyncQueue) Push(val interface{}) {
	s.q <- val
	s.size++
}

// PushWith() accepts a time.Duration, goroutine can only
// be blocked time <= duration
// Return:
// - True: push successfully
// - False: timeout
func (s *SyncQueue) PushWith(val interface{}, duration time.Duration) bool {
	select {
	case s.q <- val:
		s.size++
		return true
	case <- time.After(duration):
		return false
	}
}

func (s *SyncQueue) Pop() bool {
	select {
	case <- s.q:
		s.size--
		return true
	default:
		return false
	}
}

func (s *SyncQueue) Front() (interface{}, bool) {
	select {
	case val := <- s.q:
		return val, true
	default:
		return nil, false
	}
}

func (s *SyncQueue) Empty() bool {
	return len(s.q) == 0
}

func (s *SyncQueue) Size() int {
	return len(s.q)
}

func (s *SyncQueue) Clear() {
	for len(s.q) > 0 {
		<- s.q
	}
	s.size = 0
}

func (s *SyncQueue) Values() []any {
	values := make([]interface{}, s.size)
	for i := range values {
		values[i] = <- s.q
	}
	return values
}

func (s *SyncQueue) String() string {
	return "not impl"
}

