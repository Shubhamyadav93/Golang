package main

import "fmt"

// 1. ABSTRACTION (Contract)

type PaymentProcessor interface{
	Pay(amount float64) bool
}

// 2. IMPLEMENTATION 1: UPI

type UPI struct{
	UpiID string
}

func (u UPI) Pay(amount float64) bool{
	fmt.Println("₹ paid via UPI \n",amount,u.UpiID)
	return true
}

// 3. IMPLEMENTATION 2: Card


type Card struct {
	CardNumber string
}


func (c Card) Pay(amount float64) bool {
	fmt.Println("₹ paid via Card \n",amount,c.CardNumber)
	return true
}

// 4. APPLICATION LAYER (Interface Based Design)
// CheckoutService struct directly UPI/Card par depend nahi karta,
// yeh sirf PaymentProcessor Interface par depend karta hai.

type CheckoutService struct{
	processor PaymentProcessor 
}


func (cs CheckoutService) CompleteOrder(amount float64) {
	fmt.Println("Processing checkout...")
	success := cs.processor.Pay(amount)
	if success{
		fmt.Println("Order placed successfully! 🎉 ")
	}
}


func main()  {
	// Case A: Application UPI ke sath chalao 
	upiPayment := UPI{
		UpiID: "user@upi",
	}
	app1 := CheckoutService{
		processor: upiPayment,
	}
	app1.CompleteOrder(1500.00)

	fmt.Println("------------------------------------------")

	// Case B: Application Card ke sath chalao (Bina CheckoutService ko badle!)
	cardPayment := Card{
		CardNumber:"1234-5678-9012" ,
	}
	app2 := CheckoutService{processor: cardPayment}
	app2.CompleteOrder(3000.00)

}



// Processing checkout...
// ₹ paid via UPI 
//  1500 user@upi
// Order placed successfully! 🎉 
// ------------------------------------------
// Processing checkout...
// ₹ paid via Card 
//  3000 1234-5678-9012
// Order placed successfully! 🎉  



