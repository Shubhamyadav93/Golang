package main

import "fmt"

// 1. Interface Interface Signature (Contract)

type Payment interface{
	Pay()
}

type UPI struct{} 

func (u UPI) Pay() {
	// Complex Gateway Check, OTP verification hidden inside!
	fmt.Println("Paying via UPI...")
}

type Card struct{} 

func (c Card) Pay() {
	// Complex Gateway Check, OTP verification hidden inside!
	fmt.Println("Paying via Card...")
}

// ABSTRACTION: User ko bina internal detail dikhaye kaam karwana
// POLYMORPHISM: 'p' variable UPI ke liye alag, Card ke liye alag behave karega

func Process(p Payment) {
	p.Pay()
}


func main()  {
	Process(UPI{}) // Abstraction + Polymorphism (UPI form)
	Process(Card{}) // Abstraction + Polymorphism (Card form)
}





