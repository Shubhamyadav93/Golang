package main

import "fmt"

type Payment interface{
	Pay(amount float64) bool
}

type MockPayment struct{}

func (m MockPayment) Pay(amount float64) bool {
	fmt.Printf("Mock payment of %.0f Rs ", amount)
	return true
}

type PaymentService struct{
	payment Payment
}

func (ps PaymentService) ProcessPayment(amount float64) {
	if ps.payment.Pay(amount) {

		fmt.Println("Payment successful")
		fmt.Println("Test completed successfully")
	}
}

func main()  {
	service := PaymentService{payment: MockPayment{}}
	service.ProcessPayment(1000)
}


// Mock payment of 1000 Rs Payment successful
// Test completed successfully



