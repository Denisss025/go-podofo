package set

type Set[T comparable] struct {
	m map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{m: map[T]struct{}{}}
}

func (s *Set[T]) Put(val T) {
	s.m[val] = struct{}{}
}

func (s *Set[T]) Remove(val T) {
	delete(s.m, val)
}

func (s Set[T]) Contains(val T) bool {
	_, ok := s.m[val]

	return ok
}
