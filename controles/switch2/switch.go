package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Agora é ", t)
	switch { // switch true
	case t.Hour() < 12:
		fmt.Println("BOm dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde!")
	default:
		fmt.Println("Boa Noite!")

	}
}
