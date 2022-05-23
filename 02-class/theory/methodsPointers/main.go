package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c *Circle) setRadius(r float64) {
	c.radius = r
}

func main() {
	c := Circle{radius: 5}
	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perimeter())
	c.setRadius(10)
	fmt.Printf("%.2f\n", c.area())
	fmt.Printf("%.2f\n", c.perimeter())

}
