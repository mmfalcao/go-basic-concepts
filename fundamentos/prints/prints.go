package main

import "fmt"

func main() {
	fmt.Print("Mesma ")
	fmt.Print("Linha")

	fmt.Println(" Nova")
	fmt.Println("linha.")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("O valor de x é ", x)

	fmt.Printf("O valor de x é %.2f", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n %d, \n  %f, \n  %.1f, \n  %t, \n  %s", a, b, b, c, d)
	fmt.Printf("\n %v, \n %v, \n %v, \n  %v", a, b, c, d)

}
