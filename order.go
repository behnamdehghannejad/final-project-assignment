package domain

type Order struct {
	ID      int
	Items   []OrderItem
	Voucher Voucher
	State   OrderState
}

func NewOrder(id int) *Order {

}

func (o *Order) AddItem(product Product, qty int) error {

}

func (o *Order) ApplyVoucher(v Voucher) error {

}

func (o *Order) TotalAmount() (int, error) {

}

func (o *Order) Pay() error {

}

func (o *Order) Cancel() error {

}

func (o *Order) SnapshotItems() []OrderItem {

}
