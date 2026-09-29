package payment

import "fmt"

type Transaction struct{
	TxnID string
	Amount float64
}

func Process(t Transaction) bool {
	fmt.Printf("[PAYMENT MODULE] Successfully processed transaction %s for ₹%.2f\n",t.TxnID,t.Amount)
	return true
}

