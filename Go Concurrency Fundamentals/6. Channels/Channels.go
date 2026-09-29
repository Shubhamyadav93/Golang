package main

import (
	"fmt"
	"time"
)

// Sender Function (Gorutine A)

func sendData(ch chan string)  {
	fmt.Println("Gorutine A: Data prepare kar raha hoo...") 
	time.Sleep(2 * time.Second) // 2 Secound ka calculation time 

	fmt.Println("Gorutine A: Channel me message bhej raha hoo...")
	ch <- "Hello from Gorutine A!" // Data send Step 
	fmt.Println("Gorutine A: Data send ho gya aur Handshake complete hua!")
}



func main()  {
	
	// Step 1: String channel create kiya 
	ch:= make(chan string)

	//  Step 2: Background me Sender Gorutine launch ki 
	go sendData(ch)

	// Step 3: Main Gorutine (Gorutine B) data receive karegi 
	fmt.Println("Main Gorutine: Channel se data aane ka wait ka raha hoo...")

	msg := <- ch // Receive Step (Main Gorutine yaha Block hokar wait karegi)

	fmt.Println("Main Gorutine: Received Messag ->", msg) 

}



