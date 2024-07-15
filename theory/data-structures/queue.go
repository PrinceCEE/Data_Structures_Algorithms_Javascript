package datastructures

import (
	"errors"
)

type Queue[T any] struct {
	Data []T
}

func (q *Queue[T]) EnQueue(data T) {
	q.Data = append(q.Data, data)
}

func (q *Queue[T]) DeQueue() (T, error) {
	if q.Size() == 0 {
		var x T
		return x, errors.New("queue is empty")
	}

	data := q.Data[0]
	q.Data = q.Data[1:]

	return data, nil
}

func (q *Queue[T]) Size() int {
	return len(q.Data)
}

func (q *Queue[T]) Peek() (T, error) {
	if q.Size() == 0 {
		var x T
		return x, errors.New("queue is empty")
	}

	return q.Data[0], nil
}
