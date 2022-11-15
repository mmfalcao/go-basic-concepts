package main

import "fmt"

func media(numeros ...float64) float64 {
	total := 0.0

	for _, num := range numeros {
		total += num
	}

	return total / float64(len(numeros))
}

func PrintArrayParams(name string, params ...float64) {
	fmt.Printf("\n %s { \n", name)
	for p := range params {
		fmt.Printf("Param: %.2f \n", float64(p))
	}
	fmt.Printf("}")
}

func PrintResult(name string, result float64) {
	fmt.Printf("\n %s { \n Result: %.2f \n }", name, result)
}

func main() {
	x, y := 2.4, 5.6
	z, w, y := 7.1, 9.3, 6.2
	PrintArrayParams("Original", x, y, z, w, y)
	medXY := media(x, y)
	PrintResult("Media XY", medXY)

	medZWY := media(z, w, y)
	PrintResult("Media ZWY", medZWY)

	medBetween := media(medXY, medZWY)
	PrintResult("Media Between", medBetween)

	medAll := media(x, y, z, w, y)
	PrintResult("Media ALL", medAll)
}
