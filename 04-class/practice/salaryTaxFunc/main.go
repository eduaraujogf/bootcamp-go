package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	workedHours  int
	valuePerHour float64
}

func monthlySalaryPerHour(workedHours int, valuePerHour float64) (float64, error) {
	if workedHours < 80 {
		return 0, errors.New("the employee must not have less than 80 hours of worked hours monthly")
	}
	salary := float64(workedHours) * valuePerHour
	if salary >= 20000 {
		salary *= 0.9
	}
	return salary, nil

}
func main() {

	Employee1 := Employee{workedHours: 70, valuePerHour: 100}
	Employee2 := Employee{workedHours: 70, valuePerHour: 100}

	salary1, err := monthlySalaryPerHour(Employee1.workedHours, Employee1.valuePerHour)

	if err != nil {
		errorW := fmt.Errorf("error: %w", err)
		fmt.Println(errors.Unwrap(errorW))
	} else {
		fmt.Println(salary1)
	}

	salary2, err := monthlySalaryPerHour(Employee2.workedHours, Employee2.valuePerHour)

	if err != nil {
		errorW := fmt.Errorf("error: %w", err)
		fmt.Println(errors.Unwrap(errorW))
	} else {
		fmt.Println(salary2)
	}

}
