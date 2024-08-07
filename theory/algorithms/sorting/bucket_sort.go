package sorting

import (
	"sort"
)

func bucketSort(arr []float64) {
	n := len(arr)
	if n <= 1 {
		return
	}

	buckets := make([][]float64, n)

	// group into buckets
	for i := 0; i < n; i++ {
		index := int(arr[i] * float64(n))
		buckets[index] = append(buckets[index], arr[i])
	}

	// sort each bucket
	for i := 0; i < n; i++ {
		sort.Float64s(buckets[i])
	}

	// combine the results
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < len(buckets[i]); j++ {
			arr[index] = buckets[i][j]
			index++
		}
	}

}
