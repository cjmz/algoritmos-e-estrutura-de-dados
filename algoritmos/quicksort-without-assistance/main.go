package main

import (
	"fmt"
)

func main() {
	arr := []int{44, 67, 2, 78, 3, 15, 34}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, l int, h int) []int {
	if l < h {
		p := partition(arr, l, h)

		quicksort(arr, l, p-1)
		quicksort(arr, p+1, h)
	}
	return arr
}

func partition(arr []int, l int, h int) int {
	x := arr[h]
	i := l - 1

	for j := l; j < h-1; j++ {
		if arr[j] <= x {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	arr[i+1], arr[h] = arr[h], arr[i+1]
	return i + 1
}
