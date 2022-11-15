package main

import "fmt"

var soma = func(a, b int) int {
	return a + b
}

func main() {

	x, y := 4, 5

	fmt.Println(soma(x, y))

	sub := func(a, b int) int {
		return a - b
	}

	fmt.Println(sub(x, y))
}
