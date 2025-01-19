package stack

type Stack[T any] []T

func New[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func (s *Stack[T]) Empty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) Top() any {

	if s.Empty() {
		return nil
	}
	return (*s)[s.Size()-1]
}

func (s *Stack[T]) Push(val T) {

	*s = append(*s, val)
}

func (s *Stack[T]) Pop() any {
	if s.Empty() {
		return nil
	}

	val := (*s)[s.Size()-1]
	*s = (*s)[:s.Size()-1]
	return val
}
