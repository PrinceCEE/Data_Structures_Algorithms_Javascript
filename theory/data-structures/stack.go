package datastructures

import "errors"

type Stack[T any] struct {
	Data []T
}

func (s *Stack[T]) Add(data T) {
	s.Data = append(s.Data, data)
}

func (s *Stack[T]) Pop() (T, error) {
	size := s.Size()
	if size == 0 {
		var x T
		return x, errors.New("stack is empty")
	}

	data := s.Data[size-1]
	s.Data = s.Data[:size-1]

	return data, nil
}

func (s *Stack[T]) Size() int {
	return len(s.Data)
}

func (s *Stack[T]) Peek() (T, error) {
	size := s.Size()
	if size == 0 {
		var x T
		return x, errors.New("stack is empty")
	}

	return s.Data[size-1], nil
}
