package repository

import (
	"final-project-assignment"
)

type OrderRepository interface {
	Save(order *main.Order) error
	FindByID(id int) (*main.Order, error)
	List() []main.Order
	Delete(id int) error
	Clear()
}

type InMemoryOrderRepo struct {
	orders map[int]*main.Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{}
}

func (r *InMemoryOrderRepo) Save(order *main.Order) error {

}

func (r *InMemoryOrderRepo) FindByID(id int) (*main.Order, error) {

}

func (r *InMemoryOrderRepo) List() []main.Order {

}

func (r *InMemoryOrderRepo) Delete(id int) error {

}

func (r *InMemoryOrderRepo) Clear() {

}
