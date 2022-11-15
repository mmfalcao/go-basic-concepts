package main

import "fmt"

type curso struct {
	nome string
}

func main() {
	var coisa interface{}
	fmt.Println("Que coisa não? ", coisa)

	coisa = 3
	fmt.Println("Eita, olha essa coisa: ", coisa)

	// desse jeito ele fica um tipo mais generico
	// similar ao object do java
	type dinamico interface{}
	var coisa2 dinamico = "Opa"
	fmt.Println("Uma coisa bem dinamica: ", coisa2)

	coisa2 = true
	fmt.Println("Sera que é Uma coisa bem dinamica? ", coisa2)

	coisa2 = curso{"Golang: Explorando a Linguagem do Google"}
	fmt.Println("Qual curso vc fez coisa? ", coisa2)
}
