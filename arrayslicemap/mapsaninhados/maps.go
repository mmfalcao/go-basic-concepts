package main

import (
	"fmt"
	"strings"
)

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"J": {
			"Julia":  29.00,
			"Jiraya": 8000.00,
		},
		"L": {
			"Larissa":  3500.00,
			"Lou Reed": 5000.00,
		},
		"S": {
			"Sofia":   100.00,
			"Samurai": 7000.00,
		},
		"X": {
			"Xuxa": 7788.99,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "X")

	funcsPorLetra["V"] = map[string]float64{
		"Valetina": 7000.00,
	}

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra)
		for nome, salario := range funcs {
			fmt.Printf("%s - Salario %.2f \n", nome, salario)
		}
	}

	fmt.Println("\n --- Break --- \n")

	text := "Pelas passarelas de concreto, 	O idiota, o cidadão de bem..."
	fmt.Println(text)
	words := strings.Split(text, " ")
	occurences := map[string]int{}
	for _, w := range words {
		occurences[w]++
	}

	// count, exist := occurences["de"]
	// fmt.Println("A palavra retornou %v vezes, então é %v que ela existe", count, exist)

	for word, count := range occurences {
		fmt.Printf("Palavra %s apareceu %d vezes \n", word, count)
	}

}
