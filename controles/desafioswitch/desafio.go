package main

import "fmt"

func imprimirResultado(nota float64) {
	// if/else nao tem parenteses, apenas para uma precedencia com expressao logica
	// obrigatorio ter bloco
	if nota >= 7 {
		fmt.Println("Aprovado com nota ", nota, " Conceito: ", notaParaConceito(nota))
	} else {
		fmt.Println("Reprovado com nota ", nota, " Conceito: ", notaParaConceito(nota))
	}
}

func notaParaConceito(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}
}

func main() {

	sofia := 9.8
	valentina := 5.1
	julia := 3.2

	imprimirResultado(sofia)
	imprimirResultado(valentina)
	imprimirResultado(julia)
}
