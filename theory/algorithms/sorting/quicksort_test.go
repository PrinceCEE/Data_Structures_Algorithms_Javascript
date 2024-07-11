package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {
	arr := []int{2, 1, 0, -100, 14, 8, 19, 11}
	quickSort(arr)

	for i := range arr {
		if i == 0 {
			continue
		}

		assert.LessOrEqual(t, arr[i-1], arr[i])
	}
}
