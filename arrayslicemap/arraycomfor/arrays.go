package main

import "fmt"

func main() {
	// compilador conta
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, numero := range numeros {
		fmt.Printf("\n %d) %d\n", i, numero)
	}

	// para pegar os numeros e ignorar o indice
	for _, num := range numeros {
		fmt.Printf("\n %d\n", num)
	}

	// ignorar o indice vc vai so pegar os indices
	for num := range numeros {
		fmt.Printf("\n %d\n", num)
	}
}
