package main

import "fmt"

// 1. INTERFACE DECLARATION (Contract / Rule Set)
// Rules: Jo bhi Payment banna chahta hai, uske paas Pay() method hona chahiye.

type Payment interface{
	Pay()
}

// 2. STRUCT 1: UPI Payment

type UPI struct{
	UpiID string
}

// UPI Struct implements Payment Interface implicitly

func (u UPI) Pay() {
	fmt.Println("Paid using UPI ID:",u.UpiID)
}


// 3. STRUCT 2: Card Payment
type Card struct{
	CardNumber string
}

// Card Struct implements Payment Interface implicitly


func (c Card) Pay() {
	fmt.Println("Paid using Card Number:",c.CardNumber)
}

// 4. POLYMORPHIC FUNCTION
// Yeh function koi bhi aisa object accept kar sakta hai jo Payment interface ko satisfy karta ho.

func processPayment(p Payment)  {
	p.Pay() // // Polymorphic Call: Object ke hisab se respective Pay() method chalega

}



// 5. MAIN FUNCTION 

func main()  {
	fmt.Println("--- Go Interface & Polymorphism Demo ---")
	// Interface variable

	var p Payment 

	// Case 1: UPI Object assign kiya
	p = UPI{
		UpiID: "user@okaxis",
	}
	processPayment(p) 

	// Case 2: Card Object assign kiya
	p = Card{
		CardNumber: "4111-XXXX-XXXX-1111",
	}
	
	processPayment(p)

}


// --- Go Interface & Polymorphism Demo ---
// Paid using UPI ID: user@okaxis
// Paid using Card Number: 4111-XXXX-XXXX-1111


