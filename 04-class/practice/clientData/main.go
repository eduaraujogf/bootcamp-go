package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("customers.txt")
	if err != nil {
		fmt.Println("The file was not found or it is damaged")
		panic(err)
	}
	defer file.Close()
}
