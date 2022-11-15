package main

import "fmt"

func main() {
	x := 1
	y := 2
	fmt.Println(" Origem : ", x, y)
	// apenas pos fixado
	x++
	y--
	fmt.Println(" Postfix : ", x, y)
}
