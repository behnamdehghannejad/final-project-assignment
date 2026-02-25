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
	case Created:
		o.State = newState
	case Paid:
		if o.State == Created {
			o.State = newState
		}
	case VendorAccepted:
		if o.State == Paid {
			o.State = newState
		}
	case Shipped:
		if o.State == VendorAccepted {
			o.State = newState
		}
	case Delivered:
		if o.State == Delivered {
			o.State = newState
		}
	case Cancelled:
		o.State = newState
	default:
		return errors.New("invalid order state")
	}
	return nil
}
