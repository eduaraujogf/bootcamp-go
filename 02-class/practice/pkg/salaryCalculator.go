package pkg

const (
	C = 1000
	B = 1500
	A = 3000
)

func quantityOfHoursChecker(category, hoursQuantity float64) float64 {
	if category == B && hoursQuantity > 160 {
		return 1.20
	}
	if category == A && hoursQuantity > 160 {
		return 1.50
	}
	return 1
}
func SalaryCalculator(category, hoursQuantity float64) float64 {
	switch category {
	case C:
		return category * hoursQuantity
	case B:
		return (category * hoursQuantity) * quantityOfHoursChecker(category, hoursQuantity)
	case A:

		return (category * hoursQuantity) * quantityOfHoursChecker(category, hoursQuantity)
	}
	return 0
}
