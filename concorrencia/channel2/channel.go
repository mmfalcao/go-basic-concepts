package main

import (
	"fmt"
	"time"
)

// Channel (canal) é a forma de comunicacao entre as goroutines
// channel é um tipo da linguagem

func doisTresQuatroVezes(base int, c chan int) {
	time.Sleep(time.Second)
	c <- 2 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 3 * base // enviando dados para o canal

	time.Sleep(time.Second)
	c <- 4 * base // enviando dados para o canal
}

func main() {
	c := make(chan int)
	go doisTresQuatroVezes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c // recebendo os dados do canal
	fmt.Println("B")
	fmt.Println(a, b)
	fmt.Println(<-c)

	// fmt.Println(<-c) gera um deadlock, pois todas goroutines ja foram executadas
}
