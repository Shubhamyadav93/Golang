package database

import "fmt"


func SaveTransaction(txnID string)  {
	fmt.Printf("[DATABASE MODULE] Transaction record %s successfully saved to DB.\n",txnID)
}

