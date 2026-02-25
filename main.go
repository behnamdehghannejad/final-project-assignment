package main

import (
	"errors"
	"fmt"
)

func main() {
	repo := NewInMemoryOrderRepo()

	productId := newGenerateID(0)
	product := NewProduct(productId, "A", 100)
	product1 := NewProduct(productId, "A", 100)
	product2 := NewProduct(productId, "A", 100)
	product3 := NewProduct(productId, "A", 100)

	order := NewOrder(1)
	err := order.AddItem(product, 5)
	err = order.AddItem(product1, 5)
	err = order.AddItem(product2, 5)
	err = order.AddItem(product3, 5)
	if err == nil {
		changeOrderStatus(order, Created)
		repo.Save(order)
	}

	voucher, _ := NewFixedAmountVoucher("one", 20)
	err = order.ApplyVoucher(voucher)
	if err == nil {
		repo.Save(order)
	}

	err = order.Pay()
	if err == nil {
		changeOrderStatus(order, Paid)
	}

	fmt.Println("Vendor Accepted")

}

func changeOrderStatus(order *Order, state OrderState) {
	err := order.ChangeState(state)
	if err != nil {
		errors.New("")
	}
}

func NewProduct(generateID func() int, name string, price int) Product {
	return Product{
		ID:    generateID(),
		Name:  name,
		Price: price,
	}
}

func newGenerateID(initNumber int) func() int {
	current := initNumber
	return func() int {
		current++
		return current
	}
}
