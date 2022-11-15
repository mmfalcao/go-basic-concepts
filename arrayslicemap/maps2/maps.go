package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"Yago":   12000.00,
		"Thiago": 13000.00,
		"Abdel":  14000.00,
	}

	fmt.Println(funcsESalarios)

	funcsESalarios["Evandro"] = 9000.00

	for nome, salario := range funcsESalarios {
		fmt.Printf("%s - Salario %.2f \n", nome, salario)
	}

	// nao gera problema, se nao tiver elemento
	delete(funcsESalarios, "Inexistente")

}
