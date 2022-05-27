package main

import "fmt"

type user struct {
	name     string
	surname  string
	email    string
	products []product
}

type product struct {
	name     string
	price    float64
	quantity uint
}

func newProduct(name string, price float64) product {
	return product{name, price, 0}
}

func addProduct(u *user, p *product, quantity uint) {
	p.quantity = quantity
	u.products = append(u.products, *p)
}

func deleteProducts(u *user) {
	u.products = u.products[:0]
}

func main() {
	eduardo := user{
		name:    "Eduardo",
		surname: "Araújo",
		email:   "eduardo@mail.com",
	}
	product1 := newProduct("Laptop", 5120.2)
	product2 := newProduct("Notebook", 10.5)

	addProduct(&eduardo, &product1, 10)
	addProduct(&eduardo, &product2, 30)
	fmt.Println(eduardo)

	deleteProducts(&eduardo)
	fmt.Println(eduardo)

}
