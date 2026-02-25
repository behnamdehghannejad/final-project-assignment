package main

import (
	"errors"
	"fmt"
)

type Order struct {
	ID      int
	Items   []OrderItem
	Voucher Voucher
	State   OrderState
}

func NewOrder(id int) *Order {
	return &Order{ID: id}
}

func (o *Order) AddItem(product Product, qty int) error {
	item, err := NewOrderItem(product, qty)
	if err != nil {
		errors.New(err.Error())
	}
	o.Items = append(o.Items, item)
	return nil
}

func (o *Order) ApplyVoucher(v Voucher) error {
	amount, err := o.TotalAmount()
	if err != nil {
		errors.New("")
	}
	v.Apply(amount)
	return nil
}

func (o *Order) TotalAmount() (int, error) {
	if o == nil || o.Items == nil {
		return 0, errors.New("")
	}
	totalAmount := 0
	for _, item := range o.Items {
		totalAmount += item.TotalPrice()
	}
	return totalAmount, nil
}

func (o *Order) Pay() error {
	amount, err := o.TotalAmount()
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Printf("amount %s", amount)

	err = o.ChangeState(Paid)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

func (o *Order) Cancel() error {
	err := o.ChangeState(Cancelled)
	if err != nil {
		return errors.New(err.Error())
	}

	fmt.Println("")
	return nil
}

func (o *Order) SnapshotItems() []OrderItem {
	items := make([]OrderItem, 0, len(o.Items))
	for _, item := range o.Items {
		items = append(items, item)
	}
	return items
}
