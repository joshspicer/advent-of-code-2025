package shared

type Set[T comparable] struct {
	e map[T]struct{}
}

func CreateSet[T comparable]() *Set[T] {
	return &Set[T]{
		e: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(value T) {
	s.e[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.e, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, found := s.e[value]
	return found
}

func (s *Set[T]) Size() int {
	return len(s.e)
}

func (s *Set[T]) ForEach(fn func(T)) {
	for v := range s.e {
		fn(v)
	}
}

func (s *Set[T]) List() []T {
	keys := make([]T, 0, len(s.e))
	for key := range s.e {
		keys = append(keys, key)
	}
	return keys
}
