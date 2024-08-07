package topkfrequentelements

// Given an integer array nums and an integer k, return the k most frequent elements.
// You may return the answer in any order.
func TopKFrequentElements(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range m {
		buckets[freq] = append(buckets[freq], num)
	}

	topK := []int{}
	for i := len(buckets) - 1; i >= 0; i-- {
		for j := 0; j < len(buckets[i]); j++ {
			topK = append(topK, buckets[i][j])
			if len(topK) == k {
				return topK
			}
		}
	}

	return topK
}

// The time complexity is O(n)
// The space complexity is O(n)
