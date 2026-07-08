package main

import "fmt"

// Note: There are no inbuilt enums in Go

// Can declare custom types like this in Go
type ll int64

// Enumerated type
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Shipped
	Delivered
)

// This also works:
// type OrderStatus string

// const (
// 	Received  OrderStatus = "received"
// 	Confirmed OrderStatus = "confirmed"
// 	Shipped   OrderStatus = "shipped"
// 	Delivered OrderStatus = "delivered"
// )

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to:", status)
}

func main() {
	changeOrderStatus(Received)
}
