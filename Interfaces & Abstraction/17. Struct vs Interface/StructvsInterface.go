package main

import "fmt"

// 1. INTERFACE: Defines Behavior ("What it CAN DO")

type Payment interface{
	Pay()
}

// 2. STRUCT: Defines Data ("What it HAS")

type UPI struct{
	Mobile string
	Amount float64
}

// 3. METHOD: Connecting Data with Behavior

func (u UPI) Pay() {
	fmt.Println("Paying ₹ using Mobile: via UPI!",u.Amount,u.Mobile)
}

func main()  {
   // Step A: Struct me Data set kiya (What it HAS)
   
   myPayment := UPI{
	Mobile: "9876543210",
	Amount: 1200.0,
   }

    // Step B: Interface dwara Behavior call kiya (What it CAN DO)
   var p Payment = myPayment
   p.Pay()
}

// Paying ₹ using Mobile: via UPI! 1200 9876543210


