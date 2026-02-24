package main

import "errors"

type OrderRepository interface {
	Save(order *Order) error
	FindByID(id int) (*Order, error)
	List() []Order
	Delete(id int) error
	Clear()
}

type InMemoryOrderRepo struct {
	orders map[int]*Order
}

func NewInMemoryOrderRepo() *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		orders: make(map[int]*Order),
	}
}

func (r *InMemoryOrderRepo) Save(order *Order) error {
	if _, ok := r.orders[order.ID]; ok {
		return errors.New("Order already exists")
	}
	r.orders[order.ID] = order
	return nil
}

func (r *InMemoryOrderRepo) FindByID(id int) (*Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, errors.New("Order not found")
	}
	return order, nil
}

func (r *InMemoryOrderRepo) List() []Order {
	orders := make([]Order, len(r.orders))
	for _, order := range orders {
		orders = append(orders, order)
	}
	return orders
}

func (r *InMemoryOrderRepo) Delete(id int) error {
	if _, ok := r.orders[id]; !ok {
		return errors.New("Order not found")
	}
	delete(r.orders, id)
	return nil
}

func (r *InMemoryOrderRepo) Clear() {
	r.orders = make(map[int]*Order)
}
