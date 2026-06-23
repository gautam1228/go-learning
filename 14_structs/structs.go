package main

import (
	"fmt"
	"time"
)

// Embedded struct
type customer struct {
	id    uint8
	name  string
	phone string
}

// Order struct
type order struct {
	id        uint8 // 255 unique values
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision

	customer // Defining the embedded struct

	// Note: embedded struct can be defined like this too
	// customer struct {
	// 	name  string
	// 	phone string
	// }
}

/*
 * ----- Note -----
 * receiver type
 * this function is now attached to the struct
 * " it's a convention to use the first word of the struct
 * as the name of the var of the struct we are passing
 * into the method "
 * ----- Note -----
 */
// func (o order) changeOrderStatus(status string) {
// 	o.status = status
// }

func (o *order) changeOrderStatusWithPtr(status string) {
	o.status = status
}

// We don't pass the struct var as a ptr because this works just as well
func (o order) getOrderAmount() float32 {
	return o.amount
}

// A kinda constructor
func newOrder(id uint8, amount float32, status string, c customer) *order {
	return &order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
		customer:  c,
	}
}

func main() {
	// var order1 order = {}
	// Note: If we don't pass in any value then default values (zero value) of that var type are assigned
	// int => 0, float => 0, string => "", bool => false
	firstOrder := order{
		id:     1,
		amount: 399.99,
		status: "SHIPPED",
	}

	firstOrder.createdAt = time.Now()

	fmt.Println(firstOrder)

	// firstOrder.changeOrderStatus("DELIVERED") // Doesn't change the status because the struct instsance is passed by value
	firstOrder.changeOrderStatusWithPtr("DELIVERED")
	fmt.Println(firstOrder.getOrderAmount())
	fmt.Println(firstOrder)

	secondOrder := order{
		id:     2,
		status: "RECEIVED",
	} // here amount = 0

	fmt.Println(secondOrder)

	// In-line struct, one time creation

	language := struct {
		name   string
		isGood bool
	}{"go", true}

	fmt.Println(language)

	// Embedded struct example
	thirdOrder := order{
		id:     3,
		amount: 799.99,
		status: "RECEIVED",
		customer: customer{
			id:    1,
			name:  "John",
			phone: "1234567890",
		},
	}

	fmt.Println(thirdOrder)

	thirdOrder.customer.name = "Ron"
	fmt.Println(thirdOrder)

	fourthOrder := newOrder(
		4,
		699.99,
		"CANCELLED",
		customer{
			id:    2,
			name:  "KAREN",
			phone: "1231231231",
		},
	)
	fmt.Println(fourthOrder)
}
