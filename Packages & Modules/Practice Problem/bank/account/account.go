package account

import "fmt"

type Account struct{
	AccountNumber int 
	HolderName string
	Balance float64
}


// Unexported function (Private to bank package)
func calculateInterest(balance float64) float64  {
	return balance * 0.05 // 5% interest
}

// Exported function (Public)

func DisplayAccount(a Account)  {
	interest := calculateInterest(a.Balance)
	fmt.Println("Account Holder:",a.HolderName)
	fmt.Println("Balance:",a.Balance)
	fmt.Println("Calculated Interest:",interest)
}
