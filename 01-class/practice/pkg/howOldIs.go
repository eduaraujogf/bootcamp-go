package pkg

import "fmt"

func HowOldIs() {
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var counter int

	fmt.Println("Benjamin is", employees["Benjamin"], "years old.")

	for _, employeeAge := range employees {
		if employeeAge > 21 {
			counter++
		}
	}
	fmt.Println("There are", counter, "employees with more than 21 years old.")
	employees["Federico"] = 25
	fmt.Println("Employee Federico has been added")
	fmt.Println(employees)
	delete(employees, "Pedro")
	fmt.Println("Employee Pedro has been deleted.")
	fmt.Println(employees)
}
