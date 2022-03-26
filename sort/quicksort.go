package sort

import "github.com/ljg-cqu/gotypes"

func QuickSort[T gotypes.Number](arr []T) {
	if len(arr) < 2 {
		return
	}

	ub := len(arr) - 1
	pivot := arr[ub]
	i, j := -1, ub
	for i < j {
		for i++; arr[i] < pivot; i++ {
		}
		for j--; 0 < j && arr[j] > pivot; j-- {
		}
		arr[i], arr[j] = arr[j], arr[i]
	}

	arr[i], arr[j] = arr[j], arr[i]
	arr[i], arr[ub] = arr[ub], arr[i]

	QuickSort(arr[:i])
	QuickSort(arr[i+1:])
}
