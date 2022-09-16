package utils

type Set[T comparable] map[T]struct{}

func (s *Set[T]) Add(item T) {
	(*s)[item] = struct{}{}
}

func (s *Set[T]) Contains(item T) bool {
	if _, ok := (*s)[item]; ok {
		return true
	}
	return false
}

func (s *Set[T]) AddAll(slice []T) {
	for _, v := range slice {
		(*s).Add(v)
	}
}

func (s *Set[T]) RetainAll(slice []T) {
	for _, v := range MapKeys(*s) {
		if !Contains(slice, v) {
			delete(*s, v)
		}
	}
}
