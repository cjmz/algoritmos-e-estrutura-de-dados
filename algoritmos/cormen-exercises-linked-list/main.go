//
// Alguns exercícios do Livro Algoritmos Teoria e Prátima do Cormen
// 1 - Implementar um Insert e um Delete em uma Lista Ligada
// 2 - Implemente uma pilha usando PUSH e POP
//

package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	// If the linked list is empty, insert at LL`s head
	if l.head == nil {
		l.head = newNode
		fmt.Printf("New node added as latest node: %d\n", data)
		return
	}

	current := l.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	newNode.prev = current
	l.tail = newNode

	fmt.Printf("New node added as latest node: %d\n", data)
}

func (l *LinkedList) DeleteByValue(data int) {
	if l.head == nil {
		fmt.Println("Linked List is empty!")
		return
	}

	if l.head.data == data {
		l.head = l.head.next
		fmt.Printf("Node with value %d found and removed from linked list.\n", data)
		return
	}

	current := l.head
	prev := &Node{}

	for current != nil {
		if current.data == data {
			prev.next = current.next
			fmt.Printf("Node with value %d found and removed from linked list.\n", data)
			return
		}

		prev = current
		current = current.next
	}

	fmt.Printf("Node with value %d not found.\n", data)
}

func (l *LinkedList) Push(data int) {
	l.Insert(data)
}

func (l *LinkedList) Pop() {
	d := l.tail
	l.tail.prev.next = nil
	fmt.Printf("Node popped %d\n", d.data)
}

func main() {
	l := LinkedList{}

	l.Insert(12)
	l.Insert(36)

	l.DeleteByValue(12)
	l.DeleteByValue(38)

	l.Push(90)
	l.Push(80)

	l.Pop()
}
