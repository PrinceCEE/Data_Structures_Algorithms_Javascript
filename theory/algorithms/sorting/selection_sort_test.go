package sorting

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{2, 1, 0, -100, 14, 8, 19, 11}
	sortedArr := selectionSort(arr)

	for i := range sortedArr {
		if sortedArr[i] != arr[i] {
			t.Errorf("value in sortedArr[%d] != unsortedArr[%d]", i, i)
		}
	}
}
