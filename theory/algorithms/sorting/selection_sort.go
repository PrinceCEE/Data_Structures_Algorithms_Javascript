package sorting

// SelectionSort is a sorting algorithm that for each iteration through
// the array, it finds the smallest number and places it in its right position.
// It starts by selecting the first element as the smallest, then loops through
// the remaining array, and compares the values to know which is the smallest.
// When it finds the smallest element at the end of an iteration, it swaps the
// value with the value at the index where the iteration started, and increases the
// the index from which the next iteration will start. In SelectionSort, there is always
// a comparison.
//
// The worst case scenario is when the array is sorted in the reverse order, because for
// each iteration, there will always be a swap, and only one swap. The time complexity of
// the worst case scenario is O(n^2) because the inner loop runs for n - 1 + n - 2 +...+ 2 + 1 times.
//
// The best case scenario is when the algorithm is sorted, and there's no swaps.
//
// In the worst case scenario, SelectionSort is more performant than BubbleSort, but BubbleSort is
// better in best case scenario with a time complexity of O(n)
func selectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n; i++ {
		smallest := arr[i]
		smallestIndex := i

		for j := i; j < n; j++ {
			if arr[j] < smallest {
				smallest = arr[j]
				smallestIndex = j
			}
		}

		if smallestIndex != i {
			arr[i], arr[smallestIndex] = arr[smallestIndex], arr[i]
		}
	}

	return arr
}
