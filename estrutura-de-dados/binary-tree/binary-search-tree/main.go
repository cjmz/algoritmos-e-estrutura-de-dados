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

// Tree Insert
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

// Transplant any node to any another node
//
// This function will work as an helper to move nodes when we want to
// delete any node from our binary search tree
//
// U: It is the node that should be removed
// V: It is the node that will substitute the U node
func (t *BinaryTree) Transplant(u *Node, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

// Tree Delete
//
// # This function will remove a node from a tree
//
// Z: Node that should be removed
func (t *BinaryTree) Remove(z *Node) {
	if z.Left == nil {
		t.Transplant(z, z.Right)
	} else if z.Right == nil {
		t.Transplant(z, z.Left)
	} else {
		m := t.TreeMinimun()

		if m.Parent != z {
			t.Transplant(m, m.Right)
			m.Right = z.Right
			m.Right.Parent = m
		}

		t.Transplant(z, m)
		m.Left = z.Left
		m.Left.Parent = m
	}
}

// TreeMinimun
//
// Search the minimun value from a binary search tree
func (t *BinaryTree) TreeMinimun() *Node {
	if t.Root == nil {
		fmt.Println("Tree empty")
		return &Node{}
	}

	current := t.Root

	for current.Left != nil {
		current = current.Left
	}

	return current
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
	t.Insert(2)
	t.Insert(10)

	InorderTreeWalk(t.Root)

	t.Transplant(t.Root.Right, t.Root.Right.Right)

	fmt.Printf("Transplant Execution!\n")

	InorderTreeWalk(t.Root)

	fmt.Printf("Minumun tree value: %d\n", t.TreeMinimun().Key)

	fmt.Printf("Remove value: \n")

	t.Remove(t.Root.Right.Left)

	InorderTreeWalk(t.Root)
}
