package main

import (
	"fmt"
	"time"
)

// Normal function to be executed as a Goroutine

func printNumber() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}


func main()  {
	// 'go' keyword starts printNumbers() in a new background Goroutine

	go printNumber()
	   // Pausing Main Goroutine to allow background Goroutine to complete

	time.Sleep(time.Second)
	fmt.Println("Main Function")

}

