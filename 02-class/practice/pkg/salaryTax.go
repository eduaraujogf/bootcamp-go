package pkg

func SalaryTax(salary float64) float64 {
	const taxOfSeventeenPercent, taxOfTwentySevenPercent = 0.17, 0.27
	if salary <= 50000.00 {
		return 0.0
	} else if salary <= 150000.00 {
		return salary * taxOfSeventeenPercent
	}

	return salary * taxOfTwentySevenPercent
}
