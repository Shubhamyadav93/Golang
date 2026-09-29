package main

import (
	"errors"
	"fmt"
)

// Lower-level bank account domain function
func withdraw(balance, amount int) (int, error) {
	if amount <= 0 {
		return balance,errors.New("invalid withdrawal amount")
	}
	if amount > balance{
		return balance,errors.New("insufficient balance")
	}
	return balance - amount, nil 
}

// Service / Orchestration layer function wrapping lower-level error 

func processATM(balance,amount int)  {
	newBal , err := withdraw(balance,amount)
	if err != nil {
		// %w wraps the inner 'insufficient balance' error 
		wrappedErr := fmt.Errorf("Withdrawal failed: %w",err)
		fmt.Println(wrappedErr)
		return
	}
	fmt.Println("Withdrawal successful. New balance:",newBal)
}

func main()  {
	// Processing withdrawal request exceding balance (1500 > 1000)
	processATM(1000,1500)
}

// Withdrawal failed: insufficient balance



