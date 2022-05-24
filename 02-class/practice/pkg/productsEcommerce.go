package pkg

import "math"

const (
	Small  = "Small"
	Medium = "Medium"
	Large  = "Large"
)

type store struct {
	Products []Product
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
	Add(product Product)
}

func (p *product) costCalculator() float64 {
	switch p.Type {
	case Small:
		return 0.0
	case Medium:
		return p.Price * 1.03
	case Large:
		return p.Price*1.06 + 2500.0
	}
	return 0
}

func NewProduct(Type, Name string, Price float64) Product {
	return &product{Type, Name, Price}
}

func (s *store) Total() (message string, total float64) {
	for _, product := range s.Products {
		total += product.costCalculator()
	}
	return "Total price with all coasts: R$", math.Round(total*100) / 100
}

func (s *store) Add(product Product) {
	s.Products = append(s.Products, product)
}

func NewStore() Ecommerce {
	return &store{}
}
