package main

import "fmt"

const (
	Sum            = "+"
	Subtraction    = "-"
	Multiplication = "*"
	Division       = "/"
)

func variableVerifier(number int) {
	if number < 0 {
		fmt.Println("The number is negative")
	} else if number > 0 {
		fmt.Println("The number is positive")
	} else {
		fmt.Println("The number is zero")
	}
}

func sumFunc(value1, value2 float64) float64 {
	return value1 + value2
}

func arithmeticOperation(value1, value2 float64, operator string) float64 {
	switch operator {
	case Sum:
		return value1 + value2
	case Subtraction:
		return value1 - value2
	case Multiplication:
		return value1 * value2
	case Division:
		if value1 != 0 {
			return value1 / value2
		}
	}
	return 0
}

func main() {
	a, b, c, d := 1, 0, 5, -3
	variableVerifier(a)
	variableVerifier(b)
	variableVerifier(c)
	variableVerifier(d)

	s := sumFunc(4, 5)
	fmt.Println(s)

	fmt.Println(arithmeticOperation(6, 2, Sum))
	fmt.Println(arithmeticOperation(6, 2, Subtraction))
	fmt.Println(arithmeticOperation(6, 2, Multiplication))
	fmt.Println(arithmeticOperation(6, 2, Division))
}
