package main

import (
	"testing"
)

func TestBinarySearchIterative(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 6
	expected := true

	result := binary_search_iterative(nums, target)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 8
	expected := true

	result := binary_search_recursive(nums, target, len(nums)-1, 0)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
