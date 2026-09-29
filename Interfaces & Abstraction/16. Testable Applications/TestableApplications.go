package main

import "fmt"

// 1. COMMON INTERFACE (Contract)

type Payment interface{
	Pay()
}

// 2. PRODUCTION IMPLEMENTATION (Real UPI)

type UPI struct{}

func (u UPI) Pay() {
	   // Real Bank Server call (Asli paise katenge)
	fmt.Println("💸 Real UPI Payment Processed via Bank Server!")
}

// 3. TESTING IMPLEMENTATION (Mock Payment)

type MockPayment struct{}

func (m MockPayment) Pay() {
	// Fake success return karega (Kuch asli kaam nahi hoga)
	fmt.Println("🧪 MOCK Payment Success: Testing completed safely!")
}

// 4. APPLICATION LOGIC (Don't care if Real or Mock)

func ProcessCheckout(p Payment)  {
	fmt.Println("Order Checkout Started...")
	p.Pay() // Interface call!
}

func main()  {
	fmt.Println("--- 1. RUNNING PRODUCTION CODE ---")
   // Real Application me UPI bhejenge:
	realUPI := UPI{}
	ProcessCheckout(realUPI)

	fmt.Println("\n--- 2. RUNNING UNIT TEST SUITE ---")

	      // Unit Test likhte waqt Mock bhejenge:

	mock := MockPayment{}
	ProcessCheckout(mock)
	

}


// --- 1. RUNNING PRODUCTION CODE ---
// Order Checkout Started...
// 💸 Real UPI Payment Processed via Bank Server!

// --- 2. RUNNING UNIT TEST SUITE ---
// Order Checkout Started...
// 🧪 MOCK Payment Success: Testing completed safely!



