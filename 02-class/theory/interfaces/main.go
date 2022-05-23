package main

import (
	"fmt"
	"math"
)

const (
	rectangleType = "rectangle"
	circleType    = "circle"
)

func newFigure(geoType string, values ...float64) geometry {
	switch geoType {
	case rectangleType:
		return rectangle{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// func details(c circle) {
// 	fmt.Println(c)
// 	fmt.Println(c.area())
// 	fmt.Println(c.perimeter())
// }

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func newCircle(values float64) circle {
	return circle{radius: values}
}

func main() {
	r := rectangle{width: 3, height: 4}
	c := circle{radius: 5}
	details(r)
	details(c)
	cn := newCircle(3)
	fmt.Println(cn.area())
	fmt.Println(cn.perimeter())

	ti := newFigure(rectangleType, 2, 3)
	fmt.Println(ti.area())
	fmt.Println(ti.perimeter())
	ci := newFigure(circleType, 2)
	fmt.Println(ci.area())
	fmt.Println(ci.perimeter())

}
