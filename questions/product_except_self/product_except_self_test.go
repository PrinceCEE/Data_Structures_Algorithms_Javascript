package productexceptself

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	output := []int{24, 12, 8, 6}
	for i, v := range ProductExceptSelf([]int{1, 2, 3, 4}) {
		assert.Equal(t, output[i], v)
	}

	output = []int{0, 0, 9, 0, 0}
	for i, v := range ProductExceptSelf([]int{-1, 1, 0, -3, 3}) {
		assert.Equal(t, output[i], v)
	}
}
