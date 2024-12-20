package types

// Set is a simple set implementation.
type Set[K comparable] map[K]struct{}

// NewSet creates a set from given items.
func NewSet[K comparable](items ...K) Set[K] {
	s := make(Set[K], len(items))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// Add adds an item to the set.
func (s Set[K]) Add(item K) {
	s[item] = struct{}{}
}

// Remove removes an item from the set.
func (s Set[K]) Remove(item K) {
	delete(s, item)
}

// Contains returns true if item exists in the set.
func (s Set[K]) Contains(item K) bool {
	_, ok := s[item]
	return ok
}

// Difference returns a new set with the items that are in this set but not in other.
func (s Set[K]) Difference(other Set[K]) Set[K] {
	r := NewSet[K]()
	for item := range s {
		if !other.Contains(item) {
			r.Add(item)
		}
	}
	return r
}

// Union returns a new set with all items in this set and the other set.
func (s Set[K]) Union(other Set[K]) Set[K] {
	r := make(Set[K], max(len(s), len(other)))
	for item := range s {
		r.Add(item)
	}
	for item := range other {
		r.Add(item)
	}
	return r
}

// Intersect returns a new set with items which are common between this set and the other set.
func (s Set[K]) Intersect(other Set[K]) Set[K] {
	r := NewSet[K]()
	for item := range s {
		if other.Contains(item) {
			r.Add(item)
		}
	}
	return r
}

// Slice takes a copy of the items in the set and returns it as a slice.
// The order of items is random.
func (s Set[K]) Slice() []K {
	sl := make([]K, 0, len(s))
	for item := range s {
		sl = append(sl, item)
	}
	return sl
}
