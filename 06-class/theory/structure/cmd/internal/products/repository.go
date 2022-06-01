package products

var ps []Product
var lastID int

type repository struct{}

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, typee string, count int, price float64) (Product, error)
	LastID() (int, error)
}

func (repository) GetAll() ([]Product, error) {
	return ps, nil
}

func (repository) LastID() (int, error) {
	return lastID, nil
}

func (repository) Store(id int, name, typee string, count int, price float64) (Product, error) {
	p := Product{id, name, typee, count, price}
	ps = append(ps, p)
	lastID = p.ID
	return p, nil
}
