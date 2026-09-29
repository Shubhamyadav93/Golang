package main

import "fmt"

// 1. Common Interface (Contract)

type Payment interface{
	Pay()
}

// 2. Concrete Types

type UPI struct{}

func (u UPI) Pay() {
	fmt.Println("Payment done via UPI ✅ ")
}

type Card struct{} 

func (c Card) Pay() {
	fmt.Println("Payment done via Card ✅")
}

// 3. PaymentService - Dependency is injected via Field

type PaymentService struct{
	payment Payment  // 👈 Dependency Interface (Bahar se aayegi)

}

func (p PaymentService) Process() {
	p.payment.Pay()  // Khud create nahi kar raha, mili hui dependency use kar raha hai!
}

func main()  {
	   // --- CASE 1: Injecting UPI Dependency ---
	upiObj := UPI{}
	service1 := PaymentService{
		payment:upiObj, // 👈 Outside Injection!
	}
	service1.Process()

	   // --- CASE 2: Injecting Card Dependency ---

	cardObj := Card{}
	service2 := PaymentService{payment: cardObj}// 👈 Outside Injection!
	service2.Process()

}
