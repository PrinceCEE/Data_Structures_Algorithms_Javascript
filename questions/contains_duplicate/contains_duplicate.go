package containsduplicate

// Given an integer array nums, return true if any value appears at least twice in the array,
// and return false if every element is distinct.
func ContainsDuplicate(nums []int) bool {
	numMap := make(map[int]bool)

	for _, num := range nums {
		_, ok := numMap[num]
		if ok {
			return true
		}
		numMap[num] = true
	}

	return false
}

// Time complexity is O(n)
// Space complexity is O(n)
