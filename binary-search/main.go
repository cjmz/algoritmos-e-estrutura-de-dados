package main

import "fmt"

// Necessáriamente o array de entrada precisa estar ordenado
// para que a busca binária funcione

// A busca binária é um algoritmo de busca que encontra o
// índice de um elemento em um array ordenado em tempo
// logarítmico O(log n)

// A busca binária funciona dividindo o array em duas partes
// e verificando se o elemento que estamos procurando é maior
// ou menor que o elemento encontrado no meio do array

// Se for maior, a busca binária é chamada novamente, mas com
// o array sendo a segunda metade do array original
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(binary_search_recursive(nums, 3, len(nums)-1, 0))
	fmt.Println(binary_search_iterative(nums, 32))
}

func binary_search_recursive(nums []int, n int, hi int, lo int) bool {
	// Identifica o meio do array
	m := lo + (hi-lo)/2

	// Verfica se lower é menor que o higher
	// Isso significa que ainda há elementos para serem verificados
	if lo <= hi {
		// Verficia se o elemento encontrado é o elemento que estamos procurando
		if nums[m] == n {
			// Se for, retorna true
			return true
			// Se não for, verifica se o elemento encontrado é maior
			// que o elemento que estamos procurando
		} else if nums[m] > n {
			// Se for, chama a função novamente, mas com o higher sendo o meio - 1
			return binary_search_recursive(nums, n, m-1, 0)
			// Se não for, verifica se o elemento encontrado é menor
			// que o elemento que estamos procurando
		} else if nums[m] < n {
			// Se for, chama a função novamente, mas com o lower sendo o meio + 1
			return binary_search_recursive(nums, n, hi, m+1)
		}
	}

	// Se não encontrar o elemento, retorna false
	return false
}

// A busca binária iterativa funciona da mesma forma que a recursiva
// mas sem a necessidade de chamar a função novamente
func binary_search_iterative(nums []int, n int) bool {
	// Identifica o lower e o higher
	lo := 0
	hi := len(nums) - 1

	// Enquanto o lower for menor ou igual ao higher
	for lo <= hi {
		// Identifica o meio do array
		m := lo + (hi-lo)/2

		// Verficia se o elemento encontrado é o elemento que estamos procurando
		if nums[m] == n {
			// Se for, retorna true
			return true
			// Se não for, verifica se o elemento encontrado é maior
			// que o elemento que estamos procurando
		} else if nums[m] > n {
			// Se for, o higher passa a ser o meio - 1
			hi = m - 1
			// Se não for, verifica se o elemento encontrado é menor
			// que o elemento que estamos procurando
		} else if nums[m] < n {
			// Se for, o lower passa a ser o meio + 1
			lo = m + 1
		}
	}

	// Se não encontrar o elemento, retorna false
	return false
}
