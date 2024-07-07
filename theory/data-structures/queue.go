package datastructures

import "errors"

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) EnQueue(data T) {
	q.data = append(q.data, data)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.Size() == 0 {
		var x T
		return x, errors.New("queue is empty")
	}

	data := q.data[0]
	q.data = q.data[1:]

	return data, nil
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Size() == 0 {
		var x T
		return x, errors.New("queue is empty")
	}

	return q.data[0], nil
}
