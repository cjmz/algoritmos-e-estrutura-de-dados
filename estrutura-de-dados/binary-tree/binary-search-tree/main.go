package main

import (
	"fmt"
)

type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BinaryTree struct {
	Root *Node
}

// insertRecusive inserts a new node with the given data into the binary tree recursively.
// If the binary tree is empty, the new node becomes the root.
// Otherwise, the insert method is called on the root node to find the appropriate position for the new node.
// The parent node is updated accordingly, and the new node is inserted as the left or right child of the parent node.
func (t *BinaryTree) InsertRecursive(data int) *BinaryTree {
	if t.Root == nil {
		t.Root = &Node{Key: data}
		fmt.Printf("Value inserted to Binary Tree: %d \n", data)

	} else {
		t.Root.InsertNode(data)
	}

	return t
}

// insert inserts a new node with the given data into the binary tree.
// The method starts at the root node and traverses the tree until it finds the appropriate position for the new node.
// The parent node is updated accordingly, and the new node is inserted as the left or right child of the parent node.
func (t *BinaryTree) Insert(data int) {
	newNode := &Node{Key: data}
	var parent *Node
	current := t.Root

	for current != nil {
		parent = current

		if newNode.Key >= current.Key {
			current = current.Right
		} else {
			current = current.Left
		}
	}

	newNode.Parent = parent

	if parent == nil {
		t.Root = newNode
	} else if newNode.Key < parent.Key {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	fmt.Printf("Value inserted to Binary Tree: %d \n", data)

}

// insert is a helper method that inserts a new node with the given data into the binary tree rooted at the current node.
// If the current node is nil, the method returns without making any changes.
// If the data is less than or equal to the current node's key, the method recursively calls insert on the left child.
// If the data is greater than the current node's key, the method recursively calls insert on the right child.
func (n *Node) InsertNode(data int) {
	if n == nil {
		return
	}

	if data <= n.Key {
		if n.Left == nil {
			n.Left = &Node{Key: data}
		} else {
			n.Left.InsertNode(data)
		}
	}

	if data > n.Key {
		if n.Right == nil {
			n.Right = &Node{Key: data}
		} else {
			n.Right.InsertNode(data)
		}
	}

	fmt.Printf("Value inserted to Binary Tree: %d \n", data)
}

func InorderTreeWalk(n *Node) {
	if n != nil {
		InorderTreeWalk(n.Left)
		fmt.Printf("Node founded: %d\n", n.Key)
		InorderTreeWalk(n.Right)
	}
}

func main() {
	t := BinaryTree{}
	t.Insert(4)
	t.Insert(9)
	t.Insert(8)
	t.Insert(10)

	InorderTreeWalk(t.Root)
}
