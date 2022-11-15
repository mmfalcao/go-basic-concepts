package main

import "fmt"

type nota float64

type rangeConceito struct {
	min float64
	max float64
}

type notaConceito struct {
	conceito string
	rangeConceito
}

/*func (nc notaConceito) setConceito(conceito string, r rangeConceito) notaConceito {
	nc.conceito = conceito
	nc.rangeConceito = r
}

func getRangeMinPorConceito(n *notaConceito) float64 {
	return n.min
}

func getRangeMaxPorConceito(n *notaConceito) float64 {
	return n.max
}*/

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim
}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.99) {
		return "B"
	} else if n.entre(5.0, 6.99) {
		return "C"
	} else if n.entre(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

/*func notaParaConceito(n nota) string {
	switch n {
	case n.entre(9.0, 10.0):
		return "A"
	case n.entre(7.0, 8.99):
		return "B"
	case n.entre(5.0, 6.99):
		return "C"
	case n.entre(3.0, 4.99):
		return "D"
	case n.entre(0.0, 2.99):
		return "E"
	default:
		return "Nota inválida"
	}
}

func createConceito() (notaConceito, notaConceito, notaConceito, notaConceito, notaConceito) {
	a := notaConceito{}
	a.conceito = "A"
	a.min = 9.0
	a.max = 10.0
	fmt.Println("Conceito cadastrado: ", a)

	b := notaConceito{}
	b.conceito = "B"
	b.min = 7.0
	b.max = 8.99
	fmt.Println("Conceito cadastrado: ", b)

	c := notaConceito{}
	c.conceito = "C"
	c.min = 5.0
	c.max = 7.99
	fmt.Println("Conceito cadastrado: ", c)

	d := notaConceito{}
	d.conceito = "D"
	d.min = 3.0
	d.max = 4.99
	fmt.Println("Conceito cadastrado: ", d)

	e := notaConceito{}
	e.conceito = "E"
	e.min = 0.0
	e.max = 2.99
	fmt.Println("Conceito cadastrado: ", e)

	return a, b, c, d, e
}*/

func main() {
	// a, b, c, d, e := createConceito()
	// fmt.Println("Todos os conceitos: ", a, b, c, d, e)

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7))
	fmt.Println(notaParaConceito(5.1))
	fmt.Println(notaParaConceito(3.8))
	fmt.Println(notaParaConceito(1.2))
	fmt.Println(notaParaConceito(100))
	fmt.Println(notaParaConceito(-11))
}
