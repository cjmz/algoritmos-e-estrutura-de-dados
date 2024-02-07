package main

import (
	"fmt"
)

// Algoritmo estudado no capítulo sobre recursão do livro do Jen Wengrow
func main() {
	arr := []int{1, 2, 3, 4, 5}

	fmt.Println(recursive_array_sum(arr))
}

func recursive_array_sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + recursive_array_sum(arr[1:])
}
