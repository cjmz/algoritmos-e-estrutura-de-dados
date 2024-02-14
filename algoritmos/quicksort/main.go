package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 36, 15, 98, 14, 85}

	quicksort(arr, 0, len(arr)-1)

	fmt.Println(arr)
}

func quicksort(arr []int, low int, high int) []int {
	if low < high {
		pi := partition(arr, low, high)

		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}

	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	li := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			li++
			arr[j], arr[li] = arr[li], arr[j]
		}
	}

	arr[li+1], arr[high] = arr[high], arr[li+1]

	return li + 1
}

// 5,36,15,98,14,85 - 0 - 5
// pivot: 85
// j: 4
// li: 3
