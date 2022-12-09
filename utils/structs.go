package utils

type Set[T comparable] struct {
	vals map[T]bool
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{vals: make(map[T]bool)}
}

func (s *Set[T]) Add(val T) {
	_, ok := s.vals[val]
	if !ok {
		s.vals[val] = true
	}
}

func (s *Set[T]) Del(val T) {
	_, ok := s.vals[val]
	if ok {
		delete(s.vals, val)
	}
}

func (s *Set[T]) Length() int {
	return len(s.vals)
}
