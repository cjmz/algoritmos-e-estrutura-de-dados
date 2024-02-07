package main

import "fmt"

func main() {
	s := "abcde"

	fmt.Println(recursive_reverse_string(s))
}

func recursive_reverse_string(s string) string {
	if len(s) == 1 {
		return string(s[0])
	}

	return recursive_reverse_string(s[1:]) + string(s[0])
}
