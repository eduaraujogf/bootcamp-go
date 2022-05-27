package main

import (
	"fmt"
)

type product struct {
	name     string
	price    float64
	quantity uint
}

type service struct {
	name          string
	price         float64
	workedMinutes uint
}

type maintenance struct {
	name  string
	price float64
}

func productsSum(products []product, out chan<- float64) {
	total := 0.0
	for _, value := range products {
		total += value.price * float64(value.quantity)
	}
	out <- total
	close(out)
}

func servicesSum(services []service, out chan<- float64) {
	total := 0.0
	for _, value := range services {
		if value.workedMinutes < 30 {
			value.workedMinutes = 30
			total += value.price * float64(value.workedMinutes)
		} else {
			total += value.price * float64(value.workedMinutes)
		}
	}
	out <- total
	close(out)
}

func maintenancesSum(maintenances []maintenance, out chan<- float64) {
	total := 0.0
	for _, value := range maintenances {
		total += value.price
	}
	out <- total
	close(out)
}

func main() {
	chServices := make(chan float64)
	chProducts := make(chan float64)
	chMaintenance := make(chan float64)
	products := []product{
		{
			name:     "tires",
			price:    150.0,
			quantity: 4,
		},
		{
			name:     "lantern",
			price:    299.1,
			quantity: 10,
		},
	}
	services := []service{
		{
			name:          "replace tires",
			price:         333.0,
			workedMinutes: 100,
		},
		{
			name:          "replace lanterns",
			price:         100.52,
			workedMinutes: 20,
		},
	}
	maintenances := []maintenance{
		{
			name:  "fixed tires",
			price: 132.53,
		},
		{
			name:  "fixed lanterns",
			price: 523.14,
		},
	}

	go servicesSum(services, chServices)
	go productsSum(products, chProducts)
	go maintenancesSum(maintenances, chMaintenance)
	s1 := <-chServices
	p1 := <-chProducts
	m1 := <-chMaintenance

	fmt.Println("Services:", s1)
	fmt.Println("Products:", p1)
	fmt.Println("Maintenance:", m1)

}
