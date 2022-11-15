package main

import "fmt"

func inc1(n int) {
	n++
}

// revisao: um ponteiro é presentado por um *
func inc2(n *int) {
	// revisao: * é usado para desreferenciar, ou seja,
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) // por valor
	fmt.Println(n)

	// revisao: & usado para obter o endereco da variavel
	inc2(&n) // por referencia
	fmt.Println(n)
}
