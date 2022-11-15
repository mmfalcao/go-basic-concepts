package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma => ", a+b)
	fmt.Println("Sub => ", a-b)
	fmt.Println("div => ", a/b)
	fmt.Println("Multiplicacao => ", a*b)
	fmt.Println("Modulo => ", a%b)

	// bitwise - valor binario
	fmt.Println("E => ", a&b)   // 11 & 10 = 10
	fmt.Println("Ou => ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operacores usando math
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponenciacao => ", math.Pow(c, d))

	// logaritmo
	res := math.Log(4)
	fmt.Println("log(4) => ", res)
	res = math.Log(10.2)
	fmt.Println("log(10.2) => ", res)
	res = math.Log(-10)
	fmt.Println("log(-10) => ", res)

	// binary expoent logaritmo
	resB := math.Logb(4)
	fmt.Println("Binario Exp log(4) => ", resB)
	resB = math.Logb(10.2)
	fmt.Println("Binario Exp log(10.2) => ", resB)
	resB = math.Logb(-10)
	fmt.Println("Binario Exp log(-10) => ", resB)

	// binary logaritmo
	res2 := math.Log2(4)
	fmt.Println("2 log(4) => ", res2)
	res2 = math.Log2(10.2)
	fmt.Println("2 log(10.2) => ", res2)
	res2 = math.Log2(-10)
	fmt.Println("2 log(-10) => ", res2)

	// Decimal logaritmo
	resD := math.Log10(100)
	fmt.Println("decimal log(100) => ", resD)
	resD = math.Log10(10)
	fmt.Println("decimal log(10) => ", resD)
	resD = math.Log10(-10)
	fmt.Println("decimal log(-10) => ", resD)
}
