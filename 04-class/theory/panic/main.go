package main

import (
	"fmt"
	"os"
)

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, "It does woof woof")
}

func main() {
	fmt.Println("Starting...")
	animals := []string{
		"cow",
		"dog",
		"eager",
	}

	s := &Dog{"Sammy"}
	s = nil
	s.WoofWoof()

	fmt.Println("The animal thats fly is :", animals[len(animals)])
	_, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("End")
}
