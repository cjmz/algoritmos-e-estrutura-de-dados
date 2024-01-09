package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// Test case 1
	arr1 := []int{4, 2, 1, 3}
	expected1 := []int{1, 2, 3, 4}
	result1 := selectionSort(arr1)
	if !reflect.DeepEqual(result1, expected1) {
		t.Errorf("Expected %v, but got %v", expected1, result1)
	}

	// Test case 2
	arr2 := []int{9, 5, 7, 1, 3}
	expected2 := []int{1, 3, 5, 7, 9}
	result2 := selectionSort(arr2)
	if !reflect.DeepEqual(result2, expected2) {
		t.Errorf("Expected %v, but got %v", expected2, result2)
	}

	// Add more test cases here...
}
