package arraystack

import (
	"sync"

	"github.com/454270186/GoDazzle/list/arraylist"
)

type ArrayStack struct {
	arr *arraylist.ArrayList
	sync.Mutex
}

func New() *ArrayStack  {
	return &ArrayStack{
		arr: arraylist.New(),
	}
}

func (s *ArrayStack) Push(val interface{}) {
	s.Lock()
	s.arr.Add(val)
	s.Unlock()
}

func (s *ArrayStack) Pop() {
	s.Lock()
	s.arr.Remove(s.arr.Size()-1)
	s.Unlock()
}

func (s *ArrayStack) Top() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()

	if s.arr.Empty() {
		return nil, false
	}

	val, ok := s.arr.Get(s.arr.Size()-1)
	if !ok {
		return nil, false
	}

	return val, true
}

func (s *ArrayStack) Empty() bool {
	s.Lock()
	defer s.Unlock()
	return s.arr.Empty()
}

func (s *ArrayStack) Size() int {
	s.Lock()
	defer s.Unlock()
	return s.arr.Size()
}

func (s *ArrayStack) Clear() {
	s.Lock()
	defer s.Unlock()
	s.arr.Clear()
}

func (s *ArrayStack) Values() []any {
	s.Lock()
	defer s.Unlock()
	return s.arr.Values()
}

func (s *ArrayStack) String() string {
	return "not impl"
}

