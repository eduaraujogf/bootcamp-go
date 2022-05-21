package pkg

import "errors"

func AverageCalculator(grades ...int) (float64, error) {
	result := 0
	for _, grade := range grades {
		if grade < 0 {
			return 0, errors.New("it is not possible to enter with a negative grade")
		}
		result += grade
	}
	return float64(result / (len(grades))), nil
}
