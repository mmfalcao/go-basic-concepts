package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	fmt.Println("Inicio: ", i, p)
	// p agora recebe o endereco de memoria de i
	p = &i
	fmt.Println(" Referencia: ", i, p)
	// desreferencia o endereço, para pegar o valor
	*p++
	i++
	// Go nao tem, aritmetica de ponteiros
	// p++
	fmt.Println("DesReferenciacao: ", i, *p, p, &i)
}
