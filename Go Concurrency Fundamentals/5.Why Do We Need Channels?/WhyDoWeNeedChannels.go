package main

import "fmt"

func calculateResult(ch chan int)  {
	// Step 1: Calculation karna 
	result := 42 

	// Step 2: Channel me result bhejna (Send)
	ch <- result 
}

func main()  {

	// Step 3: Integer type ka Channel create karna 
	ch:=make(chan int)

	// Step 4: Background Gorutine launch karna 
	go calculateResult(ch) 

	// Step 5: Channel se data receive karna (Receive)
	// Main Gorutine yahan tab tak BLOCK rahega jab tak data nahi aata 
	data := <- ch 
	fmt.Println("Result received from Gorutine:", data) 


}

// Result received from Gorutine: 42 




