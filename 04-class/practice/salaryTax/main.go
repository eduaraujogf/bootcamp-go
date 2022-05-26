package main

import "fmt"

type SalaryError struct {
	msg string
}

func (e *SalaryError) Error() string {
	return fmt.Sprintf("Error: %s", e.msg)
}

func salaryChecker(salary float64) (string, error) {
	if salary < 15000 {
		return "", &SalaryError{"Salary must be higher than 15000"}
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
