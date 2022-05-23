package pkg

import "math"

const (
	Small  = "Small"
	Medium = "Medium"
	Large  = "Large"
)

type store struct {
	Products []product
}

type product struct {
	Type  string
	Name  string
	Price float64
}

type Product interface {
	costCalculator() float64
}

type Ecommerce interface {
	Total() (string, float64)
	Add(product product)
}

func (p *product) costCalculator() float64 {
	switch p.Type {
	case Small:
		return 0.0
	case Medium:
		return p.Price * 0.03
	case Large:
		return p.Price*0.06 + 2500.0
	}
	return 0
}

func NewProduct(Type, Name string, Price float64) product {
	return product{Type: Type, Name: Name, Price: Price}
}

func (s *store) Total() (message string, total float64) {
	for _, product := range s.Products {
		total += product.Price + product.costCalculator()
	}

	message = "Total price with all coasts: R$"
	return message, math.Round(total*100) / 100
}

func (s *store) Add(product product) {
	s.Products = append(s.Products, product)
}

func NewStore() Ecommerce {
	return &store{}
}
