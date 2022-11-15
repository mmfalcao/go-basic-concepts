package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))
	// uint8 unit16 unit32 uint64

	// sem sinal
	var b byte = 255
	fmt.Println("O byte é ", reflect.TypeOf(b))

	// com sinal
	i1 := math.MaxInt64
	fmt.Println("O valor maximo do int é ", i1)
	fmt.Println("o tipo de i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' // representa um mapeamento da tabelka Unicode (int32)
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2)

	// numeros reais (float32, float64)
	var x float32 = 49.99
	fmt.Println("O tipo de x é", reflect.TypeOf(x))
	fmt.Println("o tipo literal 49.99 é ", reflect.TypeOf(49.99))

	// boolean
	bo := true
	fmt.Println("O tipo de bo é ", reflect.TypeOf(bo))
	fmt.Println(bo)

	s1 := "Ola meu nome é"
	fmt.Println(s1 + " Marcel!")
	fmt.Println("o tamanho da string é ", len(s1))

	// string com multiplas linhas
	s2 := `Olá
	meu
	nome
	é`
	fmt.Println(s2 + " Marcel!")
	fmt.Println("o tamanho da string é ", len(s2))

	char := 'a'
	fmt.Println("O tipo de char é ", reflect.TypeOf(char))
	fmt.Println("O valor de char é ", char)

}
