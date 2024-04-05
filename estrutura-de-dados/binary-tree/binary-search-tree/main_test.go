package main

import (
	"testing"
)

func TestInsert(t *testing.T) {
	// Create a new binary tree
	tree := &BinaryTree{}

	// Test inserting data into an empty tree
	tree.Insert(5)
	if tree.Root == nil {
		t.Errorf("Expected root node to be created, got nil")
	}
	if tree.Root.Key != 5 {
		t.Errorf("Expected root node key to be 5, got %d", tree.Root.Key)
	}

	// Test inserting data into a non-empty tree
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)

	// Test the structure of the tree after insertion
	expectedKeys := []int{5, 3, 7, 2, 4, 6, 8}
	checkTreeStructure(t, tree.Root, expectedKeys)
}

func TestInsertRecursive(t *testing.T) {
	// Create a new binary tree
	tree := &BinaryTree{}

	// Test inserting data into an empty tree
	tree.InsertRecursive(5)
	if tree.Root == nil {
		t.Errorf("Expected root node to be created, got nil")
	}
	if tree.Root.Key != 5 {
		t.Errorf("Expected root node key to be 5, got %d", tree.Root.Key)
	}

	// Test inserting data into a non-empty tree
	tree.InsertRecursive(3)
	tree.InsertRecursive(7)
	tree.InsertRecursive(2)
	tree.InsertRecursive(4)
	tree.InsertRecursive(6)
	tree.InsertRecursive(8)

	// Test the structure of the tree after insertion
	expectedKeys := []int{5, 3, 7, 2, 4, 6, 8}
	checkTreeStructure(t, tree.Root, expectedKeys)
}

// checkTreeStructure is a helper function to check if the structure of the binary tree matches the expected keys.
func checkTreeStructure(t *testing.T, node *Node, expectedKeys []int) {
	if node == nil {
		return
	}

	// Check if the current node's key matches the expected key
	if len(expectedKeys) > 0 && node.Key != expectedKeys[0] {
		t.Errorf("Expected node key to be %d, got %d", expectedKeys[0], node.Key)
	}

	// Recursively check the left and right subtrees
	checkTreeStructure(t, node.Left, expectedKeys[1:])
	checkTreeStructure(t, node.Right, expectedKeys[1:])
}
