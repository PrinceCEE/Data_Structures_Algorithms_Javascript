package longestconsecutive

import "math"

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
func LongestConsecutive(nums []int) int {
	set := make(map[int]bool)

	for _, v := range nums {
		set[v] = true
	}

	longest := 0
	for _, v := range nums {
		if _, ok := set[v-1]; !ok {
			length := 0

			for {
				_, ok := set[v+length]
				if !ok {
					break
				}

				length++
			}

			max := math.Max(float64(longest), float64(length))
			longest = int(max)
		}
	}

	return longest
}
