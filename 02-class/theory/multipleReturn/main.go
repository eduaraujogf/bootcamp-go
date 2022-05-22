package main

import (
	"errors"
	"fmt"
)

func operations(value1, value2 float64) (float64, float64, float64, float64) {
	sum := value1 + value2
	subtraction := value1 - value2
	multiplication := value1 * value2
	var division float64

	if value2 != 0 {
		division = value1 / value2
	}

	return sum, subtraction, multiplication, division
}

func division(dividend, divider float64) (float64, error) {
	if divider == 0 {
		return 0, errors.New("the divider could not be zero.")
	}
	return dividend / divider, nil
}

func main() {
	s, r, m, d := operations(6, 2)

	fmt.Println("Sum:\t\t", s)
	fmt.Println("Subtraction:\t", r)
	fmt.Println("Multiplication:\t", m)
	fmt.Println("Division:\t", d)

	res, err := division(2, 0)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
