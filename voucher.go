package domain

type Voucher interface {
	Apply(total int) int
	Code() string
}

type PercentageVoucher struct {
	code    string
	percent int
}

func NewPercentageVoucher(code string, percent int) (*PercentageVoucher, error) {

}

func (v *PercentageVoucher) Apply(total int) int {

}
func (v *PercentageVoucher) Code() string {

}

type FixedAmountVoucher struct {
	code   string
	amount int
}

func NewFixedAmountVoucher(code string, amount int) (*FixedAmountVoucher, error) {

}

func (v *FixedAmountVoucher) Apply(total int) int {

}

func (v *FixedAmountVoucher) Code() string {

}
