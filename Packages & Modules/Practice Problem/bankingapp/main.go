package main

import (
	"bankingapp/account"
	"bankingapp/transaction"
)

func main()  {
	// Initialize Account struct instance 
	acc := account.Account{
		AccNo: 9001,
		Balamce: 10000,
	}
	// Pass memory address (&acc) so the transaction updates the original balance 
	transaction.ProcessDeposit(&acc,500)

	acc.CheckBalance()
}






