package domain

type Product struct {
	ID    int
	Name  string
	Price int
}

type OrderItem struct {
	Product  Product
	Quantity int
}

func NewOrderItem(product Product, qty int) (OrderItem, error)

func (oi OrderItem) TotalPrice() int
