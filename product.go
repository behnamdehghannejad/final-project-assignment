package main

import "errors"

type Product struct {
	ID    int
	Name  string
	Price int
}

type OrderItem struct {
	Product  Product
	Quantity int
}

func NewOrderItem(product Product, qty int) (OrderItem, error) {
	if qty <= 0 {
		return OrderItem{}, errors.New("qty must be greater than zero")
	}
	return OrderItem{
		Product:  product,
		Quantity: qty,
	}, nil
}

func (oi OrderItem) TotalPrice() int {
	return oi.Quantity * oi.Product.Price
}
