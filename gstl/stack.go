package gstl

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func NewStackWithCap[T any](cap int) *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0, cap),
	}
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack[T]) Len() int {
	return len(s.data)
}

func (s *Stack[T]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[T]) Top() T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	return s.data[len(s.data)-1]
}

func (s *Stack[T]) TopRef() *T {
	if len(s.data) == 0 {
		panic("stack is empty")
	}
	return &s.data[len(s.data)-1]
}
