package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxHeap(t *testing.T) {
	heap := NewHeap()

	t.Run("insert data", func(t *testing.T) {
		data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for _, v := range data {
			heap.Insert(v)
		}

		assert.Equal(t, 10, len(heap.Data))
	})

	t.Run("delete data", func(t *testing.T) {
		heap.Delete()

		assert.Equal(t, 9, heap.Data[0])
		// assert.Equal(t, 9, len(heap.Data))
	})
}
