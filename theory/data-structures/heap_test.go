package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	heap := NewMinHeap()

	t.Run("emptpy heap", func(t *testing.T) {
		_, err := heap.Peek()
		assert.Error(t, err)
	})

	t.Run("add data into heap", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, v := range data {
			heap.Add(v)
		}

		assert.Equal(t, 10, len(heap.Data))

		v, _ := heap.Peek()
		assert.Equal(t, 1, v)
	})

	t.Run("pop data from the heap", func(t *testing.T) {
		v, _ := heap.Pop()
		assert.Equal(t, 1, v)

		v, _ = heap.Peek()
		assert.Equal(t, 2, v)
		assert.Equal(t, 9, heap.Size())
	})
}
