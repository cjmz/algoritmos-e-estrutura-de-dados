package main

import (
	"testing"
)

func TestInsertRecursive(t *testing.T) {
	tree := BinaryTree{}
	tree.InsertRecursive(5)

	if tree.Root == nil {
		t.Errorf("Expected root node to be created, got nil")
	}

	if tree.Root.Key != 5 {
		t.Errorf("Expected root node key to be 5, got %d", tree.Root.Key)
	}
}

func TestInsert(t *testing.T) {
	tree := BinaryTree{}
	tree.Insert(8)

	if tree.Root == nil {
		t.Errorf("Expected root node to be created, got nil")
	}

	if tree.Root.Key != 8 {
		t.Errorf("Expected root node key to be 5, got %d", tree.Root.Key)
	}
}

func TestInsertHelper(t *testing.T) {
	node := &Node{Key: 5}
	node.Insert(3)

	if node.Left == nil {
		t.Errorf("Expected left child node to be created, got nil")
	}

	if node.Left.Key != 3 {
		t.Errorf("Expected left child node key to be 3, got %d", node.Left.Key)
	}

	node.Insert(7)

	if node.Right == nil {
		t.Errorf("Expected right child node to be created, got nil")
	}

	if node.Right.Key != 7 {
		t.Errorf("Expected right child node key to be 7, got %d", node.Right.Key)
	}
}
