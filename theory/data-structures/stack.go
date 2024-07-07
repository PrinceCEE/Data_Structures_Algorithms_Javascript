package datastructures

import "errors"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Add(data T) {
	s.data = append(s.data, data)
}

func (s *Stack[T]) Pop() (T, error) {
	size := s.Size()
	if size == 0 {
		var x T
		return x, errors.New("stack is empty")
	}

	data := s.data[size-1]
	s.data = s.data[:size-1]

	return data, nil
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Peek() (T, error) {
	size := s.Size()
	if size == 0 {
		var x T
		return x, errors.New("stack is empty")
	}

	return s.data[size-1], nil
}
