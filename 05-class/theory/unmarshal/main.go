package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Published bool   `json:"published"`
}

func main() {
	jsonData := `{"Name":"MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Name:", p.Name)
	fmt.Println("Price:", p.Price)
	fmt.Println("Published:", p.Published)
}
