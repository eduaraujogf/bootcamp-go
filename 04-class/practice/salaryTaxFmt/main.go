package main

import (
	"fmt"
)

func salaryChecker(salary int) (string, error) {
	if salary < 15000 {
		return "", fmt.Errorf("error: the minimal taxed salary is 15.000 and the informed salary is %d", salary)
	}
	return "You need to pay taxes", nil
}

func main() {
	const salary = 11000

	employee, err := salaryChecker(salary)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}
}
