// Package goset provides an implementation of a set
package goset

type empty struct{}

// New returns an empty set
func New[T comparable]() *Set[T] {
	return &Set[T]{
		items: make(map[T]empty),
	}
}

// From returns a new set initialized with the given items
func From[T comparable](items []T) *Set[T] {
	set := New[T]()
	for _, item := range items {
		set.Add(item)
	}
	return set
}

// Set implements set data structure using the built-in map as the underlying storage.
type Set[T comparable] struct {
	items map[T]empty
}

// Contains returns true if the set contains an item.
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.items[item]
	return ok
}

// Add adds an item to the set. Returns whether the item was newly inserted.
func (s *Set[T]) Add(item T) bool {
	if s.Contains(item) {
		return false
	}

	s.items[item] = empty{}
	return true
}

// Remove removes an item from the set. Returns whether the item was present in the set.
func (s *Set[T]) Remove(item T) bool {
	if !s.Contains(item) {
		return false
	}

	delete(s.items, item)
	return true
}

// Size returns the number of items in the set.
func (s *Set[T]) Size() int {
	return len(s.items)
}

// IsEmpty returns true if the set contains no items.
func (s *Set[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Clear clears the set, removing all items.
func (s *Set[T]) Clear() {
	s.items = make(map[T]empty)
}

// Range calls f sequentially for each item present in the set.
// If f returns false, range stops the iteration.
func (s *Set[T]) Range(f func(item T) bool) {
	for item := range s.items {
		if f(item) {
			break
		}
	}
}

// Slice returns a copy of set as slice. The order of items is not guaranteed.
func (s *Set[T]) Slice() []T {
	r := make([]T, 0, len(s.items))
	for item := range s.items {
		r = append(r, item)
	}
	return r
}
