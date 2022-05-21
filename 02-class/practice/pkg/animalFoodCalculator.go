package pkg

import "errors"

const (
	Dog       = "dog"
	Cat       = "cat"
	Hamster   = "hamster"
	Tarantula = "tarantula"
)

func funcDog(quantity float64) float64 {
	tenKilos := 10.0
	return quantity * tenKilos
}
func funcCat(quantity float64) float64 {
	fiveKilos := 5.0
	return quantity * fiveKilos
}
func funcHamster(quantity float64) float64 {
	twoHundredAndFiftyGramsToKilos := 0.250
	return quantity * twoHundredAndFiftyGramsToKilos
}
func funcTarantula(quantity float64) float64 {
	oneHundredFiftyGramsToKilos := 0.15
	return quantity * oneHundredFiftyGramsToKilos
}

func AnimalFoodCalculator(animal string) (func(quantity float64) float64, error) {
	switch animal {
	case Dog:
		return funcDog, nil
	case Cat:
		return funcCat, nil
	case Hamster:
		return funcHamster, nil
	case Tarantula:
		return funcTarantula, nil
	}

	return nil, errors.New("this animal does not exist in the database")
}
