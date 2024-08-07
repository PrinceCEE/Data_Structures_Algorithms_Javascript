package topkfrequentelements

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequentElement(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	topK := TopKFrequentElements(nums, 2)

	for _, v := range topK {
		assert.Equal(t, true, v == 1 || v == 2)
	}
}
