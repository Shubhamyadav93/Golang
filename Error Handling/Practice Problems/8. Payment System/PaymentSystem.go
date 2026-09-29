package main

import (
	"errors"
	"fmt"
)

// 1. Payment Interface

type Payment interface{
	Pay() error
}

// 2. UPI Struct implementing Payment interface

type UPI struct{
	Balance float64
	Amount float64
}

func (u UPI) Pay() error  {
	if u.Amount > u.Balance{
		return errors.New("insufficient UPI balance")
	}
	return nil
}

// 3. Service Layer combining Panic, Recover, and Wrapping 

func ProcessPaymentService(p Payment)  {
	// 🛡️ Application Boundary Safety Net
	defer func ()  {
		if r := recover(); r != nil{
			fmt.Println("Application Boundary Recovered:",r)
		} 
	}() 

	// Impossible internet state check --> Triger Panic 
	if p== nil{
		panic("impossible internal state: payment interface is nil")
	}

	// Processing Payment 
	err := p.Pay() 

	if err != nil{
		// Contextual Error Wrapping 
		fmt.Println(fmt.Errorf("payment processing failed: %w",err))
		return
	}
	fmt.Println("Payment completed successfully!")
}

func main()  {
	// Scenario 1: Insufficient Balance Error (Normal Error Flow)
	fmt.Println("----- test Case 1: Insufficient Balance -----")

	u := UPI{Balance: 500,Amount: 1000} 
	ProcessPaymentService(u)

	// Scenario 2: Nil Interface (Fatal Panic Recovered Flow)
	fmt.Println("\n-- Test Case 2: Nil Payment Interface ---")
	ProcessPaymentService(nil)

}

// ----- test Case 1: Insufficient Balance -----
// payment processing failed: insufficient UPI balance

// -- Test Case 2: Nil Payment Interface ---
// Application Boundary Recovered: impossible internal state: payment interface is nil


