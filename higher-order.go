package main

func FilterOrders(orders []Order, fn func(Order) bool) []Order
