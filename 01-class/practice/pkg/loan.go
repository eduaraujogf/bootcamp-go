package pkg

import "fmt"

func Loan(age uint, isMoreThanOneYear bool, salary float64) {

	switch {
	case (age > 22 && isMoreThanOneYear) && (salary > 100000):
		fmt.Println("You are eligible for the loan and you do not have to pay fees for this.")
	case age > 22 && isMoreThanOneYear:
		fmt.Println("You are eligible for the loan, but you will have to pay fees for this.")
	default:
		fmt.Println("You are not eligible for the loan.")
	}

}
