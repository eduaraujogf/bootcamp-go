package main

import (
	"fmt"

	"github.com/eduaraujogf/bootcamp-go/02-class/practice/pkg"
)

func main() {
	fmt.Println("----------------Exercise-01--------------")
	fmt.Printf("The salary with tax is R$%.2f", pkg.SalaryTax(50000))
	fmt.Printf("\nThe salary with tax is R$%.2f", pkg.SalaryTax(150000))
	fmt.Printf("\nThe salary with tax is R$%.2f", pkg.SalaryTax(150001))

	fmt.Println("\n------------------------------------------")

	fmt.Println("----------------Exercise-02--------------")
	eduardo, err := pkg.AverageCalculator(8, 9, 1, 7, 5)
	if err != nil {
		fmt.Println("Operation error:", err)
	} else {
		fmt.Println("The average grade of the student is", eduardo)
	}

	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercise-03--------------")

	fmt.Printf("Employee salary: R$%.2f", pkg.SalaryCalculator(pkg.C, 162))
	fmt.Printf("\nEmployee salary: R$%.2f", pkg.SalaryCalculator(pkg.B, 176))
	fmt.Printf("\nEmployee salary: R$%.2f", pkg.SalaryCalculator(pkg.A, 172))

	fmt.Println("------------------------------------------")

	fmt.Println("----------------Exercise-04--------------")

	minFunc, _ := pkg.Operation(pkg.Minimum)
	averageFunc, _ := pkg.Operation(pkg.Average)
	maxFunc, _ := pkg.Operation(pkg.Maximum)

	minValue, _ := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	averageValue, _ := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxValue, _ := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("The minimum value is:", minValue)
	fmt.Println("The average of the values is:", averageValue)
	fmt.Println("The maximum value is:", maxValue)

	fmt.Println("------------------------------")

	fmt.Println("----------------Exercise-05--------------")

	animalDog, _ := pkg.AnimalFoodCalculator(pkg.Dog)
	animalCat, _ := pkg.AnimalFoodCalculator(pkg.Cat)
	animalHamster, _ := pkg.AnimalFoodCalculator(pkg.Hamster)
	animalTarantula, _ := pkg.AnimalFoodCalculator(pkg.Tarantula)
	_, err = pkg.AnimalFoodCalculator("Dolphin")

	if err != nil {
		fmt.Println("Error:", err)
	}

	var amount float64
	amount += animalDog(1)
	amount += animalCat(1)
	amount += animalHamster(1)
	amount += animalTarantula(1)

	fmt.Printf("Amount of food to buy: %.2fkg\n", amount)

	fmt.Println("--------------------------------------")
}
