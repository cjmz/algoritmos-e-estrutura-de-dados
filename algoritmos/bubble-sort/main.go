package main

import "fmt"

func main() {
	una := []int{5, 8, 9, 2, 3, 4, 7, 6, 1}

	fmt.Println(bubbleSort(una))
}

// O BubbleSort é um algoritmo de ordenação simples.
// Ele percorre o array diversas vezes, comparando os elementos adjacentes e trocando-os de posição se o primeiro elemento for maior que o segundo.
// O processo é repetido até que não haja mais trocas.
// O BubbleSort é um algoritmo de ordenação estável, ou seja, elementos iguais não são trocados de posição.
// O BubbleSort é um algoritmo de ordenação in-place, ou seja, não é necessário alocar memória para a ordenação.
// O BubbleSort é um algoritmo de ordenação quadrático, ou seja, sua complexidade é O(n²).
func bubbleSort(arr []int) []int {
	// Percorre o array
	for i := 0; i < len(arr); i++ {
		// Percorre o array a partir do elemento atual
		for j := i + 1; j < len(arr); j++ {
			// Compara os elementos adjacentes e troca-os de posição se o primeiro elemento for maior que o segundo
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr
}
