package main

import (
	"errors"
	"fmt"
	"os"
)

type product struct {
	id       uint
	price    float64
	quantity uint
}

func (p product) lineGeneratorCSV() string {
	return fmt.Sprintf("%d, %.2f, %d\n", p.id, p.price, p.quantity)
}

func (p product) lineHeaderCSV() string {
	return "ID, Price, Quantity\n"
}

func generatorCSV(path string, products []product) error {
	if len(products) == 0 {
		return errors.New("quantity of products invalid")
	}
	file, err := os.OpenFile("./03-class/practice/fileStore/"+path, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return fmt.Errorf("error to open the file: %w", err)
	}
	defer file.Close()

	p := products[0]

	if _, err = file.WriteString(p.lineHeaderCSV()); err != nil {
		return fmt.Errorf("error to generate header line: %w", err)
	}

	for _, product := range products {
		if _, err = file.WriteString(product.lineGeneratorCSV()); err != nil {
			return fmt.Errorf("error to save line: %w", err)
		}
	}
	return nil
}

func main() {
	products := []product{
		{
			id:       5,
			price:    10.25,
			quantity: 6,
		},
		{
			id:       3,
			price:    22.25,
			quantity: 8,
		},
		{
			id:       10,
			price:    1.25,
			quantity: 23,
		},
		{
			id:       8,
			price:    332.25,
			quantity: 15,
		},
	}
	generatorCSV("products.csv", products)
}
