package utils

import "fmt"

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

type Node[T any] struct {
	Val  T
	Next *Node[T]
}

type Queue[T any] struct {
	Size int
	Head *Node[T]
	Tail *Node[T]
}

func NewQueue[T any](val T) *Queue[T] {
	q := &Queue[T]{}
	q.Enqueue(val)
	return q
}

func (q *Queue[T]) Enqueue(val T) {
	node := &Node[T]{Val: val}
	if q.Head == nil {
		q.Head = node
		q.Tail = node
		q.Size++
		return
	}

	q.Tail.Next = node
	q.Tail = node
	q.Size++
}

func (q *Queue[T]) Dequeue() T {
	val := q.Head.Val
	q.Head = q.Head.Next
	q.Size--
	return val
}

func (q *Queue[T]) Clear() {
	for q.Size != 0 {
		q.Dequeue()
	}
}

func (q *Queue[T]) Peek() T {
	return q.Head.Val
}

func QueueFromSlice[T any](slice []T) *Queue[T] {
	q := &Queue[T]{Size: 0}
	for _, val := range slice {
		q.Enqueue(val)
	}
	return q
}

func (q *Queue[T]) ToSlice() []T {
	ret := []T{}
	head := q.Head

	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

func (q *Queue[T]) DebugPrint() {
	s := q.ToSlice()
	fmt.Println(s)
}

type DefaultDict[K comparable, V any] struct {
	_map       map[K]V
	defaultVal V
}

func NewDefaultDict[K comparable, V any](defaultVal V) *DefaultDict[K, V] {
	return &DefaultDict[K, V]{
		_map:       make(map[K]V),
		defaultVal: defaultVal,
	}
}

func (d *DefaultDict[K, V]) Set(key K, val V) {
	d._map[key] = val
}

func (d *DefaultDict[K, V]) Get(key K) V {
	val, ok := d._map[key]
	if !ok {
		d._map[key] = d.defaultVal
		return d.defaultVal
	}
	return val
}
