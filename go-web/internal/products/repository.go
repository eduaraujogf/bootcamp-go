package products

import (
	"fmt"

	store "github.com/eduaraujogf/bootcamp-go/go-web/pkg"
)

//Repositorio

var ps []Product = []Product{}

var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name, typee string, count int, price float64) (Product, error)
	LastID() (int, error)
	Update(id int, name, productType string, count int, price float64) (Product, error)
	UpdateName(id int, name string) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

// func (repository) GetAll() ([]Product, error) {
// 	return ps, nil
// }
func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	r.db.Read(&ps)
	return ps, nil
}

// func (repository) LastID() (int, error) {
// 	return lastID, nil
// }

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil {
		return 0, err
	}
	if len(ps) == 0 {
		return 0, nil
	}
	return ps[len(ps)-1].ID, nil
}

// func (repository) Store(id int, name, typee string, count int, price float64) (Product, error) {
// 	p := Product{id, name, typee, count, price}
// 	ps = append(ps, p)
// 	lastID = p.ID
// 	return p, nil
// }

func (r *repository) Store(id int, name, productType string, count int, price float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, name, productType, count, price}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (repository) Update(id int, name, productType string, count int, price float64) (Product, error) {
	p := Product{Name: name, Type: productType, Count: count, Price: price}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d não encontrado", id)
	}
	return p, nil
}

func (repository) UpdateName(id int, name string) (Product, error) {
	var p Product
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("produto %d no encontrado", id)
	}
	return p, nil
}

func (repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("produto %d nao encontrado", id)
	}

	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}
