package main

import "fmt"

// 1. INTERFACE (WHAT TO DO - Façade / Exposed Layer)
// Yeh bas batata hai ki ATM kya kar sakta hai

type ATM interface{
	Withdraw(amount float64)
}

// 2. STRUCT & INTERNAL LOGIC (HOW TO DO - Hidden Implementation)
type HDFCBankATM struct{
	AccountNumber string
	Balance float64
}

// Internal Hidden Function (Aapko yeh call nahi karna padta)

func (b *HDFCBankATM) connectToBankServer() {
	fmt.Println("Connectiong to securly to HDFC Bank Server...")
}

// Actual Method Implementation
func (b *HDFCBankATM) Withdraw(amount float64) {
	b.connectToBankServer()

	if b.Balance >=amount{
		b.Balance -= amount 
		fmt.Println("Successfully withdraw ₹%2f. Remaining Balance: ₹% 2f: ", amount,b.Balance)
	}else{
		fmt.Println("Insufficient Balance!")
	}
}

func main()  {
	// User interface ke zariye ATM use kar raha hai
	var myATM ATM = &HDFCBankATM{
		AccountNumber: "ACC12345",
		Balance: 5000.0,
	}

	myATM.Withdraw(1000)
}

