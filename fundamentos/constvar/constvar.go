package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo float64 e inferido pelo compilador

	// forma reduzida de criar uma var e inicializar ela
	area := PI * m.Pow(raio, 2)
	fmt.Print("A area da circuferencia é ", area)

	const (
		a = 1
		b = 2
	)
	var (
		c = 3
		d = 4
	)
	fmt.Print("\n", a, "\n", b, "\n", c, "\n", d)

	var e, f bool = true, false
	fmt.Print("\n", e, "\n", f)
	g, h, i := 2, false, "epa!"
	fmt.Print("\n", g, "\n", h, "\n", i)
}
