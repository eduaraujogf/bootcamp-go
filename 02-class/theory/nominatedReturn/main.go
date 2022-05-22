package main

import "fmt"

func operations(value1, value2 float64) (sum, subtraction, multiplication, division float64) {
	sum = value1 + value2
	subtraction = value1 + value2
	multiplication = value1 + value2

	if value2 != 0 {
		division = value1 / value1
	}
	return sum, subtraction, multiplication, division
}

func main() {
	fmt.Println(operations(10, 20))
}
