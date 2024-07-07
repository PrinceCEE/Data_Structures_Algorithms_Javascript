package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	q := Queue[int]{}

	assert.Equal(t, 0, q.Size())

	for i := 0; i < 5; i++ {
		q.EnQueue(i + 1)
		assert.Equal(t, i+1, q.Size())
	}

	data, err := q.Peek()
	assert.NoError(t, err)
	assert.Equal(t, 1, data)

	data, err = q.DeQueue()
	assert.NoError(t, err)
	assert.Equal(t, 1, data)

	assert.Equal(t, 4, q.Size())

	size := q.Size()
	savedData := 2
	for size >= 0 {
		d, err := q.DeQueue()
		if err != nil {
			assert.Equal(t, "queue is empty", err.Error())
			break
		}

		size = q.Size()
		assert.Equal(t, savedData, d)
		savedData++
	}

	assert.Equal(t, 0, q.Size())
}
