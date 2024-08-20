package longestconsecutive

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	assert.Equal(t, 4, LongestConsecutive(nums))

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	assert.Equal(t, 9, LongestConsecutive(nums))
}
