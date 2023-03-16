package main

import (
	"fmt"
	"time"
)

// ele so vai bloquear o envio quando lotar o buffer
func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	fmt.Println("Executou")
}

func main() {
	ch := make(chan int, 3)

	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
	fmt.Println("Fim")
}
