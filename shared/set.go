package shared

type Set struct {
	e map[string]struct{}
}

func CreateSet() *Set {
	return &Set{
		e: make(map[string]struct{}),
	}
}

func (s *Set) Add(value string) {
	s.e[value] = struct{}{}
}

func (s *Set) Remove(value string) {
	delete(s.e, value)
}

func (s *Set) Contains(value string) bool {
	_, found := s.e[value]
	return found
}

func (s *Set) Size() int {
	return len(s.e)
}

func (s *Set) List() []string {
	keys := make([]string, 0, len(s.e))
	for key := range s.e {
		keys = append(keys, key)
	}
	return keys
}
