package main

import "bank/account"

func main()  {
	acc:=account.Account{
		AccountNumber: 1001,
		HolderName: "Rahul",
		Balance: 50000,
	}
	account.DisplayAccount(acc)
}

// Account Holder: Rahul
// Balance: 50000
// Calculated Interest: 2500

