package account

import "fmt"


type Account struct{
	AccNo int 
	Balamce float64 
}

// 1. Pointer Receiver to mutate (increase) balance 
func (a *Account) Deposit(amt float64) {
	a.Balamce += amt
}

// 2. Pointer Receiver with Validation Rule (decrease balance)

func (a *Account) Withdraw(amt float64) {
	if amt <= a.Balamce{
		a.Balamce -= amt
	}
}

// 3. Value Receiver (read-only balance print)
func (a Account) CheckBalance() {
	fmt.Printf("Current Balance: Rs.%2f\n",a.Balamce)
}


// Transaction: Deposited Rs.500.00
// Current Balance: Rs.10500.000000





