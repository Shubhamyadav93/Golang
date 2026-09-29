package main

import "fmt"

// Function jo sum calculate karke channel me bhejta hai
func calculate(a,b int, ch chan int)  {
	ch <- a + b // Sum calculate karke channel 'ch' me push kiya (Send operation)
}

func main()  {
	// Step 1: Ek unbuffered integer channel create kiya
	ch := make(chan int)

	// Step 2: Background Goroutine launch ki jo sum calculate karegi
	go calculate(10,20,ch)
	// Step 3: Main Goroutine channel se value aane ka wait karegi aur receive karegi
	// Note: Line `sum := <-ch` tab tak BLOCK rahegi jab tak calculate() function data send na kar de.
	sum := <- ch 

	// Step 4: Received result print kiya
	fmt.Println("Sum:",sum)
}


