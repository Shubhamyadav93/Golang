package main

import "fmt"

type Payment interface {
	Pay(amount float64)
}

type UPI struct{}

func (u UPI) Pay(amount float64) {
	fmt.Printf("Payment of %.0f Rs made through UPI ", amount)
}

type PaymentService struct{
	payment Payment
}

func NewPaymentService(p Payment) PaymentService {
	return PaymentService{payment: p}
}

func (p PaymentService) ProcessPayment(amount float64) {
	fmt.Println("Processing payment...")
	p.payment.Pay(amount)
	fmt.Println("Payment successful")
}

func main()  {
	service := NewPaymentService(UPI{})
	service.ProcessPayment(2500)
}


