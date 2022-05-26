package main

import (
	"errors"
	"fmt"
)

func salaryChecker(salary float64) (string, error) {
	if salary < 15000 {
		return "", errors.New("salary must be higher than 15000")
	}
	return "You need to pay taxes", nil
}

func main() {
	const salary = 15000

	employee, err := salaryChecker(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}
}
