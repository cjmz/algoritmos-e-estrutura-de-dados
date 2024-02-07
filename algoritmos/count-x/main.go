package main

import "fmt"

func main() {
	s := "axaxax"

	fmt.Printf("The sum of all X`s in string is: %d \n", count_x(s))
}

func count_x(s string) int {
	f := 0

	if s[0] == 120 {
		f = 1
	}

	if len(s) == 1 {
		return f
	}

	return f + count_x(s[1:])
}
