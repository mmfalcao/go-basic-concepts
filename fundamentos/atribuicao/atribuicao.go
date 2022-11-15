package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println(b)

	i := 3
	i += 3
	fmt.Println("Soma na Atribuicao =>", i)
	i -= 3
	fmt.Println("Subtrai na Atribuicao =>", i)
	i *= 2
	fmt.Println("multiplica na Atribuicao =>", i)
	i /= 2
	fmt.Println("divide na Atribuicao =>", i)
	i %= 2
	fmt.Println("modulo na Atribuicao =>", i)

	x, y := 1, 2
	fmt.Printf("\n Original: %v, %v", x, y)

	x, y = y, x
	fmt.Printf("\n inverteu: %v, %v", x, y)
}
