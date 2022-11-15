package main

import "fmt"

func obterResultado(nota float64) string {
	/* nao existe ternario
	return note >= 6 ? "Aprovado":"Reprovado"
	*/
	if nota >= 6 {
		return "Aprovado"
	}
	return "Reprovado"
}

func main() {
	fmt.Println(obterResultado(6.2))
}
