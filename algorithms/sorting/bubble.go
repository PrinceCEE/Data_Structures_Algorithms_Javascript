package sorting

// BubbleSort is a sorting algorithm that loops through
// the array n times, and for each time, compares the current
// index value with the next index value, and swaps them if the value
// at the index is greater than the value at the next index. At the end
// of each pass, the highest value is placed at its right position. Also, it can
// keep track of whether the array is sorted whenever through a pass through, there
// is no swaps.
//
// The time complexity of BubbleSort algorithm using the worst case
// scenario where the array is sorted in the reverse direction is O(n^2) because while the
// outer loop runs n times, the inner loop runs n - 1, n - 2 ... 1 for each iteration, and this
// leads to n(n - 1)/2 => O(n^2)
//
// The time complexity for the Best case scenario where the data is sorted in ascending order is O(n)
// because for the first iteration, the `sorted` value will be true and break the loop after the first
// iteration
//
// The time complexity for the Average case is still O(n^2) because there will be roughly n(n-2)/4 comparisons
// and swaps
func bubbleSort(arr []int) []int {
	currIndex := len(arr) - 1
	sorted := false

	for !sorted {
		sorted = true

		for i := 0; i < currIndex; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}

		currIndex--
	}

	return arr
}
