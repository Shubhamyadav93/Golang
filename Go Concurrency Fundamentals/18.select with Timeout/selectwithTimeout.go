package main

import (
	"fmt"
	"time"
)

func main()  {
	//  Step 1: Ek string channel create kiya data receive karna ke liye 
	ch := make(chan string)

	// Step 2: Backgound Gorutine (Server simulation)
	go func ()  {
		time.Sleep(2 * time.Second) // Server response dene me 2 secound lagata hai 
		ch <- "Response received"
	}()

	// Step 3: select block with Timeout 
	select {
	case msg:= <- ch:
		// Normal Case: Agar 1 secound ke andar data aa jata 
		fmt.Println(msg)
		
	case <- time.After(1 * time.Second):
		// Timeout Case: Agar 1 secound poora hone tak data nahi aaya 
		fmt.Println("Timeout")
	}

}

