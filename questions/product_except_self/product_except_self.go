package productexceptself

// Given an integer array nums, return an array answer such that answer[i] is equal to
// the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
func ProductExceptSelf(nums []int) []int {
	answer := []int{}

	prefix := 1
	postfix := 1

	// calculate the prefix
	for i := 0; i < len(nums); i++ {
		answer = append(answer, prefix)
		prefix *= nums[i]
	}

	// calculate the postfix
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfix
		postfix *= nums[i]
	}

	return answer
}
