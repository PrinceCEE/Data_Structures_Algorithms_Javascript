package sorting

// InsertionSort is a sorting algorithm that loops through the array from index 1,
// and compares the temporary value to the values before it. If it finds a number
// greater than the temporary value, it shifts it by 1, but if it finds a number less
// than or equal to the temporary value, it stops and goes to the next iteration
//
// The time complexity for the InsertionSort for worst case scenario(where the array is sorted in reverse)
// is O(n^2) because there are each n - 1 + n - 2 +...+ 2 + 1 comparisons and shifts (n(n-1)/2 * 2 = n^2 - 1)
//
// But for the best case scenario which is when the array is sorted, the time complexity is O(n) because for each
// iteration, the value at the index lower is already lesser than the value at the current iteration index.
//
// The time complexity for the average case scenario is O(n) as well.
func insertionSort(arr []int) []int {
	n := len(arr)

	for i := 1; i < n; i++ {
		temp := arr[i]
		position := i - 1

		for position >= 0 {
			if arr[position] > temp {
				arr[position+1] = arr[position]
				position--
			} else {
				break
			}
		}

		arr[position+1] = temp
	}

	return arr
}
