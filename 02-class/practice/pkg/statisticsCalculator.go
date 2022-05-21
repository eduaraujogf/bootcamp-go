package pkg

import "errors"

const (
	Minimum = "minimum"
	Average = "average"
	Maximum = "maximum"
)

func opMinimum(values ...float64) (float64, error) {
	min := values[0]
	for _, value := range values {
		if min > value {
			min = value
		}
	}
	return min, nil
}
func opMaximum(values ...float64) (float64, error) {
	max := values[0]
	for _, value := range values {
		if max < value {
			max = value
		}
	}
	return max, nil
}

func opAverage(values ...float64) (float64, error) {
	result := .0
	for _, value := range values {
		if value < 0 {
			return 0, nil
		}
		result += value
	}
	return result / float64(len(values)), nil
}

func Operation(operation string) (func(values ...float64) (float64, error), error) {
	switch operation {
	case Minimum:
		return opMinimum, nil
	case Maximum:
		return opMaximum, nil
	case Average:
		return opAverage, nil
	}
	return nil, errors.New("invalid function type")
}
