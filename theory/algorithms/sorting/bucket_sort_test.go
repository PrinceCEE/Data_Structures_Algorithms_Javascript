package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketSort(t *testing.T) {
	arr := []float64{0.7, 0.6, 0.1, 0.2, 0.11, 0.34, 0.5, 0.45}
	bucketSort(arr)
	assert.Equal(t, 0.1, arr[0])
	assert.Equal(t, 0.7, arr[len(arr)-1])
}
