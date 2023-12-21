package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	// Test case 1: Unsorted array
	arr1 := []int{5, 3, 8, 2, 1}
	expected1 := []int{1, 2, 3, 5, 8}
	result1 := bubbleSort(arr1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2: Sorted array
	arr2 := []int{1, 2, 3, 4, 5}
	expected2 := []int{1, 2, 3, 4, 5}
	result2 := bubbleSort(arr2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Test case 3: Empty array
	arr3 := []int{}
	expected3 := []int{}
	result3 := bubbleSort(arr3)
	if !reflect.DeepEqual(result3, expected3) {
		t.Errorf("Expected %v, but got %v", expected3, result3)
	}
}
