package main

import "fmt"

func main() {
	// var aprovados map[int]string
	// maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[7123456789] = "Julia"
	aprovados[9123456789] = "Larissa"
	aprovados[8123456789] = "Valentina"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d) \n", nome, cpf)
	}

	fmt.Println(aprovados[8123456789])
	delete(aprovados, 8123456789)
	fmt.Println(aprovados)
}
