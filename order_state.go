package main

import "errors"

type OrderState int

const (
	Created OrderState = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)

func (o *Order) ChangeState(newState OrderState) error {
	switch newState {
	case Created, Paid, VendorAccepted, Shipped, Delivered, Cancelled:
		o.State = newState
		return nil
	default:
		return errors.New("invalid order state")
	}
}
