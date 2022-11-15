package main

import "fmt"

func multiplicacao(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int {
	return funcao(p1, p2)
}

func PrintParams(name string, p1, p2 int) {
	fmt.Printf("\n %s{ \n P1: %d, \n P2 %d \n }", name, p1, p2)
}

func main() {
	x, y := 4, 5
	PrintParams("Original", x, y)

	m := multiplicacao(x, y)
	PrintParams("Multiplica", m, 0)

	resultado := exec(multiplicacao, 4, 5)
	PrintParams("Resultado", resultado, 0)
}
