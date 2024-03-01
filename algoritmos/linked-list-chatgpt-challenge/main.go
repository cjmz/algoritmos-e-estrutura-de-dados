// Lista Encadeada Simples

// 1. Implementar uma Lista Encadeada Simples: - Feito
// Criar uma estrutura de nó que contém um dado e um ponteiro para o próximo nó.
// Implementar funções para inserir nós no início e no final da lista.
// Implementar uma função para imprimir todos os valores da lista.

// 2. Deletar um Nó da Lista:
// Escreva uma função que deleta um nó com um valor específico da lista.
// A função deve ser capaz de lidar com a exclusão do primeiro ou do último nó, assim como qualquer nó intermediário.

// 3. Encontrar o n-ésimo Elemento a partir do Final:
// Escreva uma função que encontre o n-ésimo elemento a partir do final da lista.
// Implemente isso sem calcular o comprimento total da lista (tente usar o conceito de dois ponteiros).

// 4. Inverter uma Lista Encadeada:
// Escreva uma função que inverta os nós de uma lista encadeada.
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
	head *Node
	tail *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	newNode.previous = current
	l.tail = newNode
	current.next = newNode
}

func (l *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
	}

	current := l.head
	current.previous = newNode
	newNode.next = current
	l.head = newNode
}

func (l *LinkedList) DeleteByData(data int) {
	current := l.head
	previous := &Node{}

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		fmt.Printf("Succesfully deleted: %d \n", data)
		return
	}

	for current.next != nil {
		if current.data == data {
			previous.next = current.next
			fmt.Printf("Successfully deleted: %d \n", data)
		}

		previous = current
		current = current.next
	}

	if current.next == nil && current.data == data {
		previous.next = nil
		fmt.Printf("Successfully deleted: %d \n", data)
		return
	}

}

func (l *LinkedList) SelectByInverseNthElement(n int) {
	current := l.tail
	i := 0

	for current != nil {
		if i == n {
			fmt.Printf("Element %d found! Data: %d \n", n, current.data)
			return
		}

		if current.previous != nil {
			current = current.previous
		}

		i++
	}

	fmt.Printf("Element %d not found! \n", n)
}

func (l *LinkedList) Display() {
	current := l.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	for current != nil {
		fmt.Printf("Node found! Value: %d - Previous: %d - Next: %d\n", current.data, current.previous, current.next)
		current = current.next
	}

}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(30)
	list.InsertAtBeginning(5)
	// list.DeleteByData(10)
	list.InsertAtBeginning(45)
	list.Insert(105)
	// list.DeleteByData(105)

	list.SelectByInverseNthElement(2)
	list.Display()
	// fmt.Println("Hello World")
}
