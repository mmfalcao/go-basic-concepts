package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 5, 6:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Nota inválida"
	}
}

func main() {

	for k := 1; k <= 10; k++ {
		fmt.Println(notaParaConceito(float64(k)))
	}

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(3.8))
	fmt.Println(notaParaConceito(1.2))
	fmt.Println(notaParaConceito(100))
	fmt.Println(notaParaConceito(-11))
}
