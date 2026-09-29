package main

import (
	"fmt"
	
)

type Payment interface{
	Pay(amount float64) bool
}

type Notification interface{
	Send(message string)
}

type UPI struct{}

func (u UPI) Pay(amount float64) bool {
	fmt.Println("Processing UPI payment...")
	return true
}
type Email struct{}

func ( e Email) Send(msg string) {
	fmt.Println("Sending confirmation email...")
}

type OrderService struct{
	payment Payment
	notification Notification
}

func NewOrderService(p Payment,n Notification) OrderService {
	return OrderService{payment: p,notification: n}
}

func (o OrderService) PlaceOrder(customer,product string, amount float64) {
	fmt.Printf("Order placed for:%s Product:%s Amount:%.0f",customer,product,amount)
	if o.payment.Pay(amount){
		fmt.Println("Payment successfully")
		o.notification.Send("Order confirmend")
		fmt.Println("Order confirmed successfully")
	}
}

func main()  {
	orderAPP := NewOrderService(UPI{},Email{})
	orderAPP.PlaceOrder("Rahul","Laptop",50000)
}


// Order placed for:Rahul Product:Laptop Amount:50000Processing UPI payment...
// Payment successfully
// Sending confirmation email...
// Order confirmed successfully



