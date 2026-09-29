package main

import "fmt"

func main(){
	// Step 1: Buffered channel create kiya jiski capacity 3 hai 
	ch := make(chan int,3)

	// Step 2: Main Goroutine bina block hue teeno values push kar sakti hai 

	ch <- 10 
	ch <- 20 
	ch <- 30 

	// Note: Is point par buffer poora full ho chuka hai (capacity 3/3 feilled)

	// Step 3: FIFO (First In, First Out) Order me values recive karega 

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

}


