package set

type Set[T comparable] struct {
	items map[T]struct{}
}

func New[T comparable](elements ...T) *Set[T] {
	s := &Set[T]{items: make(map[T]struct{})}
	for _, element := range elements {
		s.Add(element)
	}
	return s
}

func (s *Set[T]) Add(element T) {
	s.items[element] = struct{}{}
}

func (s *Set[T]) Remove(element T) {
	delete(s.items, element)
}

func (s *Set[T]) Contains(element T) bool {
	_, exists := s.items[element]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Elements() []T {
	elements := make([]T, 0, len(s.items))
	for item := range s.items {
		elements = append(elements, item)
	}
	return elements
}
