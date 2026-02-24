package main

func FilterOrders(orders []Order, fn func(Order) bool) []Order {
	result := make([]Order, 0, len(orders))
	for _, order := range orders {
		if fn(order) {
			result = append(result, order)
		}
	}
	return result
}
