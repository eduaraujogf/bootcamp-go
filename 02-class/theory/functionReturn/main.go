package main

import "fmt"

func opSum(value1, value2 float64) float64 {
	return value1 + value2
}

func opSubtraction(value1, value2 float64) float64 {
	return value1 - value2
}

func opMultiplication(value1, value2 float64) float64 {
	return value1 * value2
}

func opDivision(value1, value2 float64) float64 {
	if value2 == 0 {
		return 0
	}
	return value1 / value2
}

func arithmeticOperation(operator string) func(value1, value2 float64) float64 {
	switch operator {
	case "Sum":
		return opSum
	case "Subtraction":
		return opSubtraction
	case "Multiplication":
		return opMultiplication
	case "Division":
		return opDivision
	}
	return nil
}

func main() {
	oper, div := arithmeticOperation("Sum"), arithmeticOperation("Division")
	r, d := oper(2, 5), div(10, 2)
	fmt.Println(r, d)
}
