package gotype

import "sync"

type Set struct {
	m map[interface{}]bool
	sync.RWMutex
}

func NewSet() *Set {
	return &Set{
		m:make(map[interface{}]bool),
	}
}

func (s *Set)Add(item interface{})  {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set)Remove(item interface{})  {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *Set)Contain(item interface{}) bool  {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set)Len() int {
	return len(s.m)
}

func (s *Set)Clear()  {
	s.RLock()
	defer s.RUnlock()
	s.m = map[interface{}]bool{}
}

func (s *Set)IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set)List() []interface{} {
	s.RLock()
	defer s.RUnlock()
	list := []interface{}{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}