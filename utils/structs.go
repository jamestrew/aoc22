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

type HashMap[K comparable, V any] struct {
	_map map[K]V
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	return &HashMap[K, V]{_map: make(map[K]V)}
}

func (h *HashMap[K, V]) Set(key K, val V) {
	_, ok := h._map[key]
	if !ok {
		h._map[key] = val
	}
}

func (h *HashMap[K, V]) Get(key K, _default V) V {
	val, ok := h._map[key]
	if !ok {
		return _default
	}
	return val
}

func (h *HashMap[K, V]) Del(key K) V {
	val, ok := h._map[key]
	if ok {
		delete(h._map, key)
	}
	return val
}
