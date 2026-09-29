package main

import "fmt"

type Payment interface{
	Pay()
}

type UPI struct{
	Mobile string
	Amount float64
}

func (u UPI) Pay() {
	fmt.Println("Making Rs payment through UPI",u.Mobile)
	fmt.Println("Mobile Number:",u.Mobile)
}


type Card struct{
	CardNumber string
	Amount float64
}

func (c Card) Pay() {
	fmt.Println("Making  Rs payment through Card", c.Amount)
	fmt.Println("card Number:", c.CardNumber)
}

func makePayment(p Payment)  {
	p.Pay()
}

func main()  {
	u:=UPI{
		Mobile: "9876543210",
		Amount: 1000,
	}
	makePayment(u)

	c := Card{
		CardNumber:"ACFH39721",
		Amount: 75397,
	}
	makePayment(c)
}

// Making Rs payment through UPI 9876543210
// Mobile Number: 9876543210
// Making  Rs payment through Card 75397
// card Number: ACFH39721



