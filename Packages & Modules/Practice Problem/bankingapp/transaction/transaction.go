package transaction

import (
	"bankingapp/account"
	"fmt"
)

// Accepts a pointer (*acount.Account) to directly modify the original objct

func ProcessDeposit(acc *account.Account,amount float64)  {
	acc.Deposit(amount)
	fmt.Printf("Transaction: Deposited Rs.%.2f\n",amount)
}



