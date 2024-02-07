package main

import "testing"

func TestRecursiveArraySum(t *testing.T) {
	// table of test cases
	tests := []struct {
		input  []int
		expect int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 2, 4, 5}, 14},
		{[]int{1, 1, 1, 1, 2}, 6},
	}

	for _, test := range tests {
		output := recursive_array_sum(test.input)

		if output != test.expect {
			t.Errorf("recursive_array_sum(%v) = %d; want %d", test.input, output, test.expect)
		}
	}

}
