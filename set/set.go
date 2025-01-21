package set

// Set представляет собой множество.
type Set struct {
	elements map[interface{}]struct{}
}

// New создает новое множество.
func New() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

// Add добавляет элемент в множество.
func (s *Set) Add(element interface{}) {
	s.elements[element] = struct{}{}
}

// Remove удаляет элемент из множества.
func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)
}

// Contains проверяет, содержится ли элемент в множестве.
func (s *Set) Contains(element interface{}) bool {
	_, exists := s.elements[element]
	return exists
}

// Size возвращает количество элементов в множестве.
func (s *Set) Size() int {
	return len(s.elements)
}

// Clear очищает множество.
func (s *Set) Clear() {
	s.elements = make(map[interface{}]struct{})
}

// Elements возвращает все элементы множества.
func (s *Set) Elements() []interface{} {
	keys := make([]interface{}, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}
