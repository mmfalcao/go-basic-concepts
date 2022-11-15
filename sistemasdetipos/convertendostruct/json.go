package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	// struct para json
	p1 := produto{1, "Notebook", 1899.90, []string{"Promoção", "Eletronico"}}
	p1ToJson, _ := json.Marshal(p1) // slice de byte, error
	fmt.Println(string(p1ToJson))

	var p2 produto
	jsonString := `{"id":2,"nome":"Mouse","preco":89.5,"tags":["Saldão","Periferico"]}`
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2.Tags[1])
}
