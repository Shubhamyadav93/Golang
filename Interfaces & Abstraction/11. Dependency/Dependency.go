package main

import "fmt"

// 1. Dependency Rules (Contract)
 
type Payment interface{
	Pay(amount float64)
}

// 2. Concrete Types (Jo kaam karenge)

type UPI struct{}

func (uu UPI) Pay(amount float64) {
	fmt.Println("₹ paid via UPI\n",amount)
}

// 3. Main Struct (Jishki dependency Payment par hai)

type PaymentService struct{
	payment Payment //👈 Dependency Field
}

// Function jo dependency ka use kar raha hai
func (ps PaymentService) ProcessOrder(amount float64) {
	fmt.Println("Starting Order Processing...")

	// PaymentService ne apni dependency ko kaam karne ke liye use kiya:

	ps.payment.Pay(amount) 

}

func main()  {
	// Step A: Pehle dependency object create kiya
	myUpi := UPI{}

	// Step B: PaymentService me uski dependency (myUpi) provide/inject ki
	service := PaymentService{
		payment: myUpi,
	}

	   // Step C: Execution
	service.ProcessOrder(500.0)
}


