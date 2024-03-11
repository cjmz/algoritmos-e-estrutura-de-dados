// Lista Encadeada Simples

// 1. Implementar uma Lista Encadeada Simples:
// Criar uma estrutura de nó que contém um dado e um ponteiro para o próximo nó. - OK
// Implementar funções para inserir nós no início e no final da lista. - ok
// Implementar uma função para imprimir todos os valores da lista. - ok

// 2. Deletar um Nó da Lista:
// Escreva uma função que deleta um nó com um valor específico da lista. - ok
// A função deve ser capaz de lidar com a exclusão do primeiro ou do último nó, assim como qualquer nó intermediário. - ok

// 3. Encontrar o n-ésimo Elemento a partir do Final:
// Escreva uma função que encontre o n-ésimo elemento a partir do final da lista. - ok
// Implemente isso sem calcular o comprimento total da lista (tente usar o conceito de dois ponteiros). - ok

// 4. Inverter uma Lista Encadeada:
// Escreva uma função que inverta os nós de uma lista encadeada. - ok
// Tente fazer isso iterativamente e, em seguida, recursivamente.
// Lista Duplamente Encadeada

// 5. Implementar uma Lista Duplamente Encadeada:
// Expandir a estrutura de nó anterior para conter um ponteiro adicional para o nó anterior.
// Implementar funções para inserir e deletar nós de uma lista duplamente encadeada.

// 6. Ordenar uma Lista Encadeada:
// Implemente uma função para ordenar uma lista encadeada (pode ser uma simples lista encadeada para começar).
// Você pode escolher qualquer algoritmo de ordenação, mas um merge sort adaptado para listas encadeadas é uma boa opção.

// 7. Remover Duplicatas de uma Lista Ordenada:
// Escreva uma função que remova todos os elementos duplicados de uma lista encadeada ordenada.

// 8. Rotação de Lista Encadeada:
// Implemente uma função que rotacione a lista encadeada para a direita por k lugares.

// Desafios Adicionais

// 9. Detectar Ciclo em uma Lista Encadeada:
// Implemente uma função que verifique se uma lista encadeada possui um ciclo (um nó que aponta de volta a um nó anterior).

// 10. Ponto de Interseção de Duas Listas Encadeadas:
// - Escreva uma função que determine se duas listas encadeadas se cruzam e, em caso afirmativo, retorne o ponto de interseção.

package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type LinkedList struct {
	first *Node
	last  *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if l.first == nil {
		l.first = newNode
		fmt.Printf("Node value %d added to the end \n", newNode.data)
		return
	}

	current := l.first

	for current.next != nil {
		current = current.next
	}

	l.last = newNode
	newNode.previous = current
	current.next = newNode
	fmt.Printf("Node value %d added to the end \n", newNode.data)
}

func (l *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}

	if l.first == nil {
		l.first = newNode
		fmt.Printf("Node value %d added to the beginning \n", newNode.data)
		return
	}

	l.first.previous = newNode
	newNode.next = l.first
	l.first = newNode
	fmt.Printf("Node value %d added to the beginning \n", newNode.data)
}

func (l *LinkedList) DeleteByValue(data int) {
	if l.first == nil {
		fmt.Println("Linked list empty!")
		return
	}

	if l.first.data == data {
		l.first = l.first.next
		fmt.Printf("Value %d found and removed from the beginning! \n", data)
		return
	}

	current := l.first.next
	previous := l.first

	for current.next != nil {
		if current.data == data {
			previous.next = current.next
			fmt.Printf("Value %d found and removed! \n", data)
			return
		}

		previous = current
		current = current.next
	}

	if current.data == data {
		previous.next = nil
		fmt.Printf("Value %d found and removed from the end \n", data)
		return
	}

	fmt.Printf("Value %d not found \n", data)

}

func (l *LinkedList) PrintList() {
	if l.first == nil {
		fmt.Println("Linked list empty!")
		return
	}

	current := l.first

	for current.next != nil {
		fmt.Printf("Node found! Value: %d | Next: %p \n", current.data, current)
		current = current.next
	}

	fmt.Printf("Last node found! Value: %d | Next: %p \n", current.data, current)
}

func (l *LinkedList) FoundByNthElement(n int) {
	current := l.last
	i := 0

	for current != nil {
		if i == n {
			fmt.Printf("Node %dNth found! Value: %d\n", n, current.data)
			return
		}

		i++
		// fmt.Printf("Node %dNth found! Value: %d\n", n, current.data)
		// fmt.Println(i)
		current = current.previous
	}

	fmt.Println("Element not found!")
}

func (l *LinkedList) Reverse() {
	if l.first == nil {
		fmt.Println("Linked list empty!")
		return
	}

	current := l.first
	var next *Node
	var prev *Node

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.first = prev

	fmt.Println("List inverted")
}

func (l *LinkedList) RecursiveReverse(current *Node) {
	if current == nil {
		return
	}

	next := current.next
	prev := current.previous

	current.next = prev

	l.RecursiveReverse(next)
}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.InsertAtBeginning(350)
	list.Insert(35)
	list.InsertAtBeginning(290)
	list.DeleteByValue(10)
	list.DeleteByValue(15)

	fmt.Println("=========================")

	list.PrintList()

	fmt.Println("=========================")

	list.FoundByNthElement(0)

	fmt.Println("=========================")

	list.Reverse()

	fmt.Println("=========================")

	list.PrintList()

	fmt.Println("=========================")

	list.RecursiveReverse(list.first)

	fmt.Println("=========================")

	list.PrintList()

	fmt.Println("=========================")

	// fmt.Println("Hello, let`s start again, it is just to review our learning! Don`t worry!")
}
