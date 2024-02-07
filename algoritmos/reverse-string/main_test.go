package main

import (
	"testing"
)

func TestRecursiveReverseString(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{"abcde", "edcba"},
		{"claudio", "oidualc"},
		{"samuka", "akumas"},
	}

	for _, test := range tests {
		output := recursive_reverse_string(test.input)

		if output != test.expect {
			t.Errorf("recursive_reverse_string(%s): %s; want %s", test.input, output, test.expect)
		}

	}
}
