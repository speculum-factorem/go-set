package set

type Set struct {
	elements map[interface{}]struct{}
}

func New() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

func (s *Set) Add(element interface{}) {
	s.elements[element] = struct{}{}
}

func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)
}

func (s *Set) Contains(element interface{}) bool {
	_, exists := s.elements[element]
	return exists
}

func (s *Set) Size() int {
	return len(s.elements)
}

func (s *Set) Clear() {
	s.elements = make(map[interface{}]struct{})
}

func (s *Set) Elements() []interface{} {
	keys := make([]interface{}, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
