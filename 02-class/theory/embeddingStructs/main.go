package main

import "fmt"

type Vehicle struct {
	km   float64
	hour float64
}

func (v Vehicle) details() {
	fmt.Printf("km:\t%.2f\nhour:\t%.2f\n", v.km, v.hour)
}

type Car struct {
	v Vehicle
}

func (c *Car) Accelerate(minutes int) {
	c.v.hour = float64(minutes) / 60
	c.v.km = c.v.hour * 100
}

func (c *Car) Details() {
	fmt.Println("\nV:\tAutomobile")
	c.v.details()
}

type Motorcycle struct {
	v Vehicle
}

func (m *Motorcycle) Accelerate(minutes int) {
	m.v.hour = float64(minutes) / 60
	m.v.km = m.v.hour * 80
}

func (m *Motorcycle) Details() {
	fmt.Println("\nV:\tMotorcycle")
	m.v.details()
}

func main() {
	car := Car{}
	car.Accelerate(360)
	car.Details()

	motorcycle := Motorcycle{}
	motorcycle.Accelerate(360)
	motorcycle.Details()

}
