package twosums

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSums(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		toPass bool
	}{
		{
			arr:    []int{2, 7, 11, 15},
			target: 9,
			toPass: true,
		},
		{
			arr:    []int{3, 2, 4},
			target: 6,
			toPass: true,
		},
		{
			arr:    []int{3, 5, 8, 11, 14},
			target: 30,
			toPass: false,
		},
		{
			arr:    []int{3, 3},
			target: 6,
			toPass: true,
		},
	}

	for _, v := range tests {
		res := TwoSums(v.arr, v.target)
		if v.toPass {
			sum := v.arr[res[0]] + v.arr[res[1]]
			assert.Equal(t, v.target, sum)
		} else {
			assert.Nil(t, res)
		}
	}
}
