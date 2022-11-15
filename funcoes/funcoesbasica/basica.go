package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2 \n P1: %s, \n P2: %s \n", p1, p2)
}

func f3() string {
	// fmt.Println("F3")
	return "F3"
}

func f4(p1 string, p2 string) string {
	return fmt.Sprintf("F4 %s %s", p1, p2)
}

func f5() (string, string) {
	return "Retorno 1", "retorno 2"
}

func main() {

	f1()
	f2("Param1", "Param2")

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Printf("Function -> 1 %s \n", r3)
	fmt.Printf("Function -> 2 %s \n", r4)

	r51, r52 := f5()

	fmt.Printf("F5 -> R1 %s \n", r51)
	fmt.Printf("F5 -> R2 %s \n", r52)
}
