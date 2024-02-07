package main

import "testing"

func TestCountX(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"axaxax", 3},
		{"xxxx", 4},
		{"xyz", 1},
	}

	for _, test := range tests {
		output := count_x(test.input)

		if output != test.expect {
			t.Errorf("count_x(%s): %d; want %d", test.input, output, test.expect)
		}
	}
}
