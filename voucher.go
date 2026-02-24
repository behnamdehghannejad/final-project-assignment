package main

import "errors"

type Voucher interface {
	Apply(total int) int
	Code() string
}

type PercentageVoucher struct {
	code    string
	percent int
}

func NewPercentageVoucher(code string, percent int) (*PercentageVoucher, error) {
	if percent < 0 || percent > 100 {
		return nil, errors.New("percentage must be between 0 and 100")
	}
	return &PercentageVoucher{code, percent}, nil
}

func (v *PercentageVoucher) Apply(total int) int {
	if v == nil {
		return total
	}

	totalAfterDiscount := total - v.percent*total
	if totalAfterDiscount < 0 {
		return total
	}
	return totalAfterDiscount
}
func (v *PercentageVoucher) Code() string {
	return v.code
}

type FixedAmountVoucher struct {
	code   string
	amount int
}

func NewFixedAmountVoucher(code string, amount int) (*FixedAmountVoucher, error) {
	if amount < 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	return &FixedAmountVoucher{code: code, amount: amount}, nil
}

func (v *FixedAmountVoucher) Apply(total int) int {
	if v == nil {
		return total
	}

	if totalAfterDiscount := total - v.amount; totalAfterDiscount < 0 {
		return total
	} else {
		return totalAfterDiscount
	}
}

func (v *FixedAmountVoucher) Code() string {
	return v.code
}
