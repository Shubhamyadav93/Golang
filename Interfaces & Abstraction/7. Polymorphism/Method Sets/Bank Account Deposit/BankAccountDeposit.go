package main

import "fmt"

type BankAccount struct{
	Owner string
	Balance int 
}

func (b BankAccount) CheckBalance() {
	fmt.Printf("Owner: %s , Balance %d",b.Owner,b.Balance)
}

func (b *BankAccount) Deposit(amount int) {
	b.Balance = amount
}

func main()  {
	var owner string 
	var amount int 

	fmt.Scan(&owner)
	fmt.Scan(&amount)

	account := &BankAccount{
		Owner: owner,
		Balance: 0,
	}

	account.Deposit(amount)
	account.CheckBalance()
}



