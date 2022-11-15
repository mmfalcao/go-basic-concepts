package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return // retorno limpo
	// return p2, p1
}

func PrintParams(name string, p1, p2 int) {
	fmt.Printf("\n %s{ \n P1: %d, \n P2 %d \n }", name, p1, p2)
}

func main() {
	primeiro := 1
	segundo := 2
	PrintParams("Original", primeiro, segundo)

	i1, i2 := trocar(primeiro, segundo)

	PrintParams("Trocados", i1, i2)
}
