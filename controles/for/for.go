package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // nao tem while em Go
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 20; j += 2 {
		fmt.Printf("\n %d ", j)
	}

	fmt.Println("\nMisturando...")
	for k := 1; k <= 10; k++ {
		if k%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("ImPar")
		}
	}

	for {
		// laço infinito
		fmt.Println("Para Sempre...")
		time.Sleep(time.Second * 5)
	}
}
