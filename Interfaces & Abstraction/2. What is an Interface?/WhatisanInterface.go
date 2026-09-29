package main

import "fmt"

// 1. INTERFACE (Contract - Sirf signature)

type Payment interface{
	Pay()
}

// 2. STRUCT 1 (UPI Payment) 
type UPI struct{
	UpiID string
}


func (u UPI) Pay() { // Satisfies Payment interface 
	fmt.Println("Paying via UPI ID:",u.UpiID) 
}


// 3. STRUCT 2 (CreditCard Payment) 

type CreditCard struct{
	CardNumber string
}

func (c CreditCard) Pay() {
	fmt.Println("Paying via Credit Card Number:",c.CardNumber)
}

// 4. MAIN FUNCTION 

func main()  {
	var p Payment 

	// UPI payment execute kar sakte hain 
	p = UPI{
		UpiID: "arun@okaxis",
	}
	p.Pay()

	// Credit Card payment execute kar sakte hain 
	p = CreditCard{
		CardNumber: "4111-XXXX-XXXX-1111",
	}
	p.Pay()
}



// Paying via UPI ID: arun@okaxis
// Paying via Credit Card Number: 4111-XXXX-XXXX-1111
