package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (b *BankAccount) deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) withdraw(amount float64) {
	if amount <= b.Balance {
		b.Balance -= amount
	} else {
		fmt.Println("Insufficient Balance")
	}
}

func (b BankAccount) displayBalance() {
	fmt.Println("%.Of",b.Balance)
}

func main() {
	acc := BankAccount{
		AccountNumber: "1001",
		HolderName:    "Arun",
		Balance:       10000,
	}

	fmt.Println("Account Holder:",acc.HolderName)
	fmt.Print("Intial Balance:  ")
	acc.displayBalance() 

	acc.deposit(5000)
	fmt.Print("After Deposit: ")
	acc.displayBalance()

	acc.withdraw(3000)
	fmt.Print("After Withdrawal: ")
	acc.displayBalance()

}


// Account Holder: Arun
// Intial Balance:  %.Of 10000
// After Deposit: %.Of 15000
// After Withdrawal: %.Of 12000



