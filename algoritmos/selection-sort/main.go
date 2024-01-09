package main

import (
	"fmt"
)

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	fmt.Println(selectionSort(arr))
}

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		tli := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[tli] {
				tli = j
			}
		}

		if tli != i {
			arr[i], arr[tli] = arr[tli], arr[i]
		}
	}

	return arr
}
