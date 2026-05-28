package stack

type Stack[T any] struct {
	items    []T
	capacity int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		items:    []T{},
		capacity: 0,
	}
}

func NewWithSize[T any](size int) *Stack[T] {
	if size < 0 {
		size = 0
	}

	return &Stack[T]{
		items:    []T{},
		capacity: size,
	}
}

func (s *Stack[T]) Push(value T) bool {
	if s.IsFull() {
		return false
	}

	s.items = append(s.items, value)
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	lastIndex := len(s.items) - 1
	value := s.items[lastIndex]
	s.items = s.items[:lastIndex]

	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T

	if s.IsEmpty() {
		return zero, false
	}

	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) IsFull() bool {
	if s.capacity == 0 {
		return false
	}

	return len(s.items) == s.capacity
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Capacity() int {
	return s.capacity
}
