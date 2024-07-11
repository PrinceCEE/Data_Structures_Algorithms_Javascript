package sorting

// QuickSort is a sorting algorithm that recursively partitions halves and sort each child array.
// It partitions the array, gets the index of the partition, which means the element in that index is
// in its right position. Then, it calls itself on the two halves of the present array ([:pivotIndex] & [pivotIndex+1:])
//
// The time complexity for the best case and average case is O(nlogn) because the time complexity of the partitioning at
// each level is O(n), and there are logn levels.
//
// The time complexity of the worst case scenario is O(n^2), and this happens when the array is sorted in ascending or
// descending order because while the time complexity of the partition is O(n), there will be n levels which results to
// O(n*n) => O(n^2)
func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivotIndex := partition(arr)
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func partition(arr []int) int {
	n := len(arr)
	prePivotIdx := -1
	pivot := arr[n-1]

	for i := 0; i < n-1; i++ {
		if arr[i] < pivot {
			prePivotIdx++
			arr[i], arr[prePivotIdx] = arr[prePivotIdx], arr[i]
		}
	}

	arr[prePivotIdx+1], arr[n-1] = arr[n-1], arr[prePivotIdx+1]
	return prePivotIdx + 1
}
