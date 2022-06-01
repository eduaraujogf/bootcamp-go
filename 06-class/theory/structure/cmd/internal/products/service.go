package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, typee string, count int, price float64) (Product, error)
}

type service struct {
	repository Repository
}

func (service) GetAll() ([]Product, error)
func (service) Store(name, typee string, count int, price float64) (Product, error)
