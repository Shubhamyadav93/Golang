package main

import (
	"fmt"
	"goprojects3/database"
	"goprojects3/payment"
	"goprojects3/validation"
	"log"
)

func main()  {
	fmt.Println("=== Starting Order Processing System ===")

	txnID := "TXN_998877"
	amount := 2500.00

	// Step 1: Use Validation Module 
	err := validation.ValidationAmount(amount)
	if err != nil{
		log.Fatalf("[ERROR] Validation Failed: %v",err)
	}

	// Step 2: Use Payment Module 
	txn := payment.Transaction{
		TxnID: txnID,
		Amount: amount,
	}
	isSuccess := payment.Process(txn)

	// Step 3: Use Database Module 
	if isSuccess{
		database.SaveTransaction(txnID)
	}

}


// [PAYMENT MODULE] Successfully processed transaction TXN_998877 for ₹2500.00
// [DATABASE MODULE] Transaction record TXN_998877 successfully saved to DB.
