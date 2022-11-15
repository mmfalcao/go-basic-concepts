package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// metodo: funcao com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}
	fmt.Println("Proximo produto: ", produto1)
	fmt.Println("Qual preco com desconto? ", produto1.precoComDesconto())

	produto2 := produto{"Notebook", 1989.90, 0.10}
	fmt.Println("Proximo produto: ", produto2)
	fmt.Println("Qual preco com desconto? ", produto2.precoComDesconto())
}
