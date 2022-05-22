package main

import "fmt"

func Variables() {
	var name string
	var hours int
	hours = 20
	product, price := "jeans", 10.5

	var (
		productVarBlock  = "Course"
		quantityVarBlock = 25
		priceVarBlock    = 40.50
		inStock          = true
	)

	const Status = "ok"

	const (
		ProductConstBlock  = "Course"
		QuantityConstBlock = 20
		PriceConstBlock    = 40.50
	)

	fmt.Println(name, hours, product, price, productVarBlock, quantityVarBlock, priceVarBlock, inStock, Status, ProductConstBlock, QuantityConstBlock, PriceConstBlock)
}
