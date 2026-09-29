package main

import (
	"fmt"
	"time"
)

// Function 1: Even numbers print karne ke liye (2, 4, 6, 8, 10)
func printEven()  {
	for i:=2;i<=10;i+=2{
		println("Even:",i)
		time.Sleep(50 * time.Millisecond)// Har iteration ke baad 50ms pause

	}
}
	// Function 2: Odd numbers print karne ke liye (1, 3, 5, 7, 9)

func printOdd()  {
	
	for i:= 1; i<=10;i+=2{
		fmt.Println("Odd:",i)
		time.Sleep(50 * time.Millisecond) // Har iteration ke baad 50ms pause 

	}

}

func main()  {
	// Step 1: Background Goroutine 1 launch ki
	go printEven()

	// Step 2: Background Goroutine 2 launch ki
	go printOdd()

	// Step 3: Main Goroutine ko 400ms tak roka
	// (Dono functions me 5 loops × 50ms = 250ms lagenge, 400ms safe buffer time hai)
	time.Sleep(400 * time.Millisecond)

}





