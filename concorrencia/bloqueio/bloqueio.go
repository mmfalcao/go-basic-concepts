package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante
	fmt.Println("So depois que foi lido")
}

func main() {
	c := make(chan int) // canal sem buffer

	go rotina(c)

	fmt.Println(<-c) // operacao bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}
