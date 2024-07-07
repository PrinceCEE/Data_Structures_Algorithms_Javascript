package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := Stack[int]{}

	assert.Equal(t, 0, s.Size())

	for i := 0; i < 5; i++ {
		s.Add(i + 1)
		assert.Equal(t, i+1, s.Size())
	}

	data, err := s.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 5, data)

	data, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 5, data)

	assert.Equal(t, 4, s.Size())

	size := s.Size()
	savedData := 4
	for size >= 0 {
		d, err := s.Pop()
		if err != nil {
			assert.Equal(t, "stack is empty", err.Error())
			break
		}

		size = s.Size()
		assert.Equal(t, savedData, d)
		savedData--
	}

	assert.Equal(t, 0, s.Size())
}
