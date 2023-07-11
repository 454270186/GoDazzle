package syncmap

import (
	"fmt"
	"hash/fnv"
	"sync"
)

const shareCount = 32

type SyncMap struct {
	maps []*syncMapShared
}

type syncMapShared struct {
	sync.RWMutex
	m map[any]any
}

func New() *SyncMap {
	sMap := &SyncMap{}
	sMap.maps = make([]*syncMapShared, shareCount)
	for i := 0; i < shareCount; i++ {
		sMap.maps[i] = &syncMapShared{m: make(map[any]any)}
	}

	return sMap
}

func (s *SyncMap) Put(key, value interface{}) {
	shareMap := s.getShare(key)
	shareMap.Lock()
	defer shareMap.Unlock()
	shareMap.m[key] = value
}

func (s *SyncMap) Get(key interface{}) (interface{}, bool) {
	shareMap := s.getShare(key)
	shareMap.RLock()
	defer shareMap.RUnlock()
	val, ok := shareMap.m[key]
	return val, ok
}

func (s *SyncMap) Remove(key interface{}) {
	shareMap := s.getShare(key)
	shareMap.Lock()
	defer shareMap.Unlock()
	delete(shareMap.m, key)
}

func (s *SyncMap) Keys() []interface{} {
	keys := []interface{}{}
	for _, share := range s.maps {
		if share.m == nil {
			continue
		}
		share.RLock()
		for k := range share.m {
			keys = append(keys, k)
		}
		share.RUnlock()
	}

	return keys
}

func (s *SyncMap) Empty() bool {
	for _, share := range s.maps {
		share.RLock()
		if len(share.m) != 0 {
			return false
		}
		share.RUnlock()
	}

	return true
}

func (s *SyncMap) Size() int {
	size := 0
	for _, share := range s.maps {
		share.RLock()
		size += len(share.m)
		share.RUnlock()
	}

	return size
}

// @WARN: Clear() is not thread safe
func (s *SyncMap) Clear() {
	s.maps = s.maps[:0]
}

func (s *SyncMap) Values() []any {
	vals := make([]any, s.Size())
	i := 0
	for _, share := range s.maps {
		share.RLock()
		if share.m == nil {
			share.RUnlock()
			continue
		}
		for _, v := range share.m {
			vals[i] = v
			i++
		}
		share.RUnlock()
	}

	return vals
}

func (s *SyncMap) String() string {
	return "not impl"
}

func (s *SyncMap) getShare(key interface{}) *syncMapShared {
	hashFunc := func(key any) uint {
		h := fnv.New32()
		h.Write([]byte(fmt.Sprintf("%v", key)))
		return uint(h.Sum32())
	}
	return s.maps[hashFunc(key)%shareCount]
}

