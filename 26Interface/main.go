package main

import (
	"fmt"
	"time"
)

type Paymenter interface {
	Pay(amount float32) // only write function signture
	Refund(amount float32)
}

// Payment  engine

type Payment struct {
	gateway Paymenter // which is passed as gateway that should have a "Pay(amount float32) method"
}

func (p Payment) MakePayment(amount float32) {
	p.gateway.Pay(amount)

}
func (p Payment) MakeRefund(amount float32) {
	p.gateway.Refund(amount)

}

//  RazorPay Gateway

type RazorPay struct{}

func (r RazorPay) Pay(amount float32) {
	fmt.Printf("Payment is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\nPayment", amount, "done by razorpay gateway.")
}
func (r RazorPay) Refund(amount float32) {
	fmt.Printf("Refund is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\nRefund", amount, "done by razorpay gateway.")
}

//  PhonePe Gateway

type PhonePe struct{}

func (p PhonePe) Pay(amount float32) {
	fmt.Printf("Payment is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\nPayment", amount, " done by phonepe gateway.")
}
func (p PhonePe) Refund(amount float32) {
	fmt.Printf("Refund is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\nRefund", amount, " done by phonepe gateway.")
}

//  Stripe Gateway

type Stripe struct{}

func (s Stripe) Pay(amount float32) {
	fmt.Printf("Payment is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\nPayment", amount, " done by stripe gateway.")
}
func (s Stripe) Refund(amount float32) {
	fmt.Printf("Refund is progreesing..")
	for i := 0; i < 3; i++ {
		time.Sleep((4 - 1) * time.Second)
		fmt.Printf("..")
	}
	time.Sleep(3 * time.Second)
	fmt.Println("\n Refund", amount, " done by stripe gateway.")
}

func main() {
	phonePeGW := PhonePe{}
	// stripeGW := Stripe{}
	// razorPayGW := RazorPay{}

	payment := Payment{
		gateway: phonePeGW,
	}
	payment.MakeRefund(200)

}
