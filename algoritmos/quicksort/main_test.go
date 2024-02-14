package main

import (
	"reflect"
	"testing"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{[]int{2, 5, 6, 4, 7}, []int{2, 4, 5, 6, 7}},
		{[]int{15, 69, 5, 98, 1}, []int{1, 5, 15, 69, 98}},
	}

	for _, test := range tests {
		quicksort(test.input, 0, len(test.input)-1)

		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("quicksort(%v); want %v", test.input, test.expect)
		}
	}
}
