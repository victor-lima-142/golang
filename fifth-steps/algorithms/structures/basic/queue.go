package basic

type Queue[T any] struct {
	data []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.data = append(q.data, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var item T
		return item, false
	}

	item := q.data[0]
	q.data = q.data[1:]

	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Peek() (T, bool) {
	if q.IsEmpty() {
		var item T
		return item, false
	}

	return q.data[0], true
}

func (q *Queue[T]) Clear() {
	q.data = make([]T, 0)
}
