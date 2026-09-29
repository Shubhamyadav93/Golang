package main

import "fmt"

// Producer Function: Orders create karke channel me bhejta hai

func producer(ch chan int)  {
	// Step 1: 101 se 105 tak Order IDs channel me push karna
	for id := 101; id<= 105; id ++{
		ch <- id 
	}
	// Step 2: Sabhi 5 orders bhejne ke baad Channel CLOSE kar diya
	close(ch)

}

func main()  {
	// Step 3: Unbuffered channel allocate kiya
	ch := make(chan int)

	// Step 4: Background Goroutine (Producer) launch ki
	go producer(ch)
	// Step 5: Main Goroutine (Processor) range loop se data consume karegi
	// Ye loop tab tak chalega jab tak channel close nahi ho jata aur khali na ho jaye
	for orderID := range ch {
		fmt.Printf("Processing Order: %d\n",orderID)
	}
	// Range loop auto-exit hone ke baad program yahan seamlessly khatam ho jayega

}
