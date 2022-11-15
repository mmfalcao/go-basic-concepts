package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	fmt.Printf("\n Lista de Aprovados \n")
	for i, aprovado := range aprovados {
		fmt.Printf("%d %s \n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Sofia", "Larissa", "Valentina", "Julia"}
	imprimirAprovados(aprovados...)

}
