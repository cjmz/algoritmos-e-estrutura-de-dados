package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"oi": 4, "tchau": 5}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m)
}
