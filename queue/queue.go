package queue

type Queue[T any] struct {
	items    []T
	capacity int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		items:    []T{},
		capacity: 0,
	}
}

func NewWithSize[T any](size int) *Queue[T] {
	if size < 0 {
		size = 0
	}

	return &Queue[T]{
		items:    []T{},
		capacity: size,
	}
}

func (q *Queue[T]) Enqueue(value T) bool {
	if q.IsFull() {
		return false
	}

	q.items = append(q.items, value)
	return true
}

func (q *Queue[T]) Dequeue() (T, bool) {
	var zero T

	if q.IsEmpty() {
		return zero, false
	}

	value := q.items[0]
	q.items = q.items[1:]

	return value, true
}

func (q *Queue[T]) Front() (T, bool) {
	var zero T

	if q.IsEmpty() {
		return zero, false
	}

	return q.items[0], true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) IsFull() bool {
	if q.capacity == 0 {
		return false
	}

	return len(q.items) == q.capacity
}

func (q *Queue[T]) Size() int {
	return len(q.items)
}

func (q *Queue[T]) Capacity() int {
	return q.capacity
}

func (q *Queue[T]) Values() []T {
	values := make([]T, len(q.items))
	copy(values, q.items)

	return values
}
