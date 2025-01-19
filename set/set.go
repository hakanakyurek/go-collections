package set

type Set[T comparable] map[T]bool

func New[T comparable]() Set[T] {
	return make(Set[T])
}

func (s *Set[T]) Add(key T) {
	(*s)[key] = true
}

func (s *Set[T]) Remove(key T) {
	delete((*s), key)
}

func (s *Set[T]) Contains(key T) bool {
	return (*s)[key]
}

func FromIter[T comparable](elements ...T) Set[T] {
	set := make(Set[T])

	for _, ele := range elements {
		set.Add(ele)
	}

	return set
}
