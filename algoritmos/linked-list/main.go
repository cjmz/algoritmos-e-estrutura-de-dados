package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
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

	current.next = newNode
}

func (l *LinkedList) Display() {
	current := l.head

	if current == nil {
		fmt.Println("Linked list empty!")
		return
	}

	fmt.Println("Starting linked list display:")

	for current != nil {
		fmt.Printf("Node value: %d | Next node: %d \n", current.data, current.next)
		current = current.next
	}

	fmt.Println("End display")

}

func main() {
	list := LinkedList{}

	list.Insert(10)
	list.Insert(36)
	list.Insert(300)
	list.Insert(48)

	list.Display()
	// fmt.Println("Hello")
}
