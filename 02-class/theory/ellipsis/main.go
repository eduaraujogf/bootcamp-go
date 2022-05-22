package main

import "fmt"

const (
	Sum            = "+"
	Subtraction    = "-"
	Multiplication = "*"
	Division       = "/"
)

func sum(values ...float64) float64 {
	var result float64
	for _, value := range values {
		result += value
	}

	return result
}

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

func arithmeticOperation(operator string, values ...float64) float64 {
	switch operator {
	case Sum:
		return operationReturn(values, opSum)
	case Subtraction:
		return operationReturn(values, opSubtraction)
	case Multiplication:
		return operationReturn(values, opMultiplication)
	case Division:
		return operationReturn(values, opDivision)
	}
	return 0
}

func operationReturn(values []float64, operation func(value1, value2 float64) float64) float64 {
	var result float64
	for i, value := range values {
		if i == 0 {
			result = value
		} else {
			result = operation(result, value)
		}
	}
	return result
}

func main() {
	fmt.Println(sum(2, 3, 4, 5, 56, 6))

	fmt.Println(arithmeticOperation(Multiplication, 2, 3, 5, 6, 7, 1, 20))
}
