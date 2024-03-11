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

func (l *LinkedList) Push(data int) {
	newNode := &Node{data: data}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		fmt.Printf("Pushed: %d\n", data)
		return
	}

	newNode.next = l.head
	l.head = newNode
	fmt.Printf("Pushed: %d\n", data)
}

func (l *LinkedList) Pop() {
	fmt.Printf("Popped: %d\n", l.head.data)
	next := l.head.next
	l.head = next
}

func main() {
	l := LinkedList{}
	l.Push(10)
	l.Push(45)
	l.Push(80)
	l.Pop()
	l.Pop()
	l.Push(75)
	l.Pop()

}
