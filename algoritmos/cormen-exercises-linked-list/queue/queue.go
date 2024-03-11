// Develop a Queue in Golang using a Simply Linked List this should happen in O(1)

package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Enqueue(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		fmt.Printf("Enqueued: %d\n", data)
		return
	}

	l.tail.next = newNode
	l.tail = newNode
	fmt.Printf("Enqueued: %d\n", data)
}

func (l *LinkedList) Dequeue() {
	if l.head == nil {
		fmt.Println("Linked list empty!")
		return
	}

	fmt.Printf("Dequeued: %d\n", l.head.data)
	next := l.head.next
	l.head = next
}

func main() {
	l := LinkedList{}
	fmt.Println("Init:")
	l.Enqueue(10)
	l.Enqueue(30)
	l.Dequeue()
	l.Enqueue(50)
	l.Dequeue()
	l.Dequeue()
	l.Dequeue()
}
