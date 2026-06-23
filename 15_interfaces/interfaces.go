package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, accountNo string)
}

type payment struct {
	gateway paymenter
}

// for any struct to implement an interface, it needs to have methods
// which match the signatures of the methods defined in the interface

// Razorpay payments
type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Printf("Paying %v via Razorpay...\n", amount)
}

func (r razorpay) refund(amount float32, accountNo string) {
	fmt.Printf("Refunding %v to Account: %v via Razorpay...\n", amount, accountNo)
}

// Stripe payments
type stripe struct{}

func (r stripe) pay(amount float32) {
	fmt.Printf("Paying %v via Stripe...\n", amount)
}

func (r stripe) refund(amount float32, accountNo string) {
	fmt.Printf("Refunding %v to Account: %v via Stripe...\n", amount, accountNo)
}

func main() {
	// Using razorpay
	razorPayPaymentGw := razorpay{}
	newPayment1 := payment{
		gateway: razorPayPaymentGw,
	}

	newPayment1.gateway.pay(100)

	// Using stripe
	stripePaymentGw := stripe{}
	newPayment2 := payment{
		gateway: stripePaymentGw,
	}

	newPayment2.gateway.pay(100)

}
