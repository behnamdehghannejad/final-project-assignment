package domain

import "final-project-assignment"

type OrderState int

const (
	Created OrderState = iota
	Paid
	VendorAccepted
	Shipped
	Delivered
	Cancelled
)

func (o *main.Order) ChangeState(newState OrderState) error {

}
