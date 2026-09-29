package main

import "fmt"

// Function jo 3 subjects ke marks accept karke total sum channel me bhejta hai
func calcTotal(m1, m2, m3 int , ch chan int)  {
	ch <- m1 + m2 + m3 // Sum (80 + 75 + 90 = 245) calculate karke channel me push kiya
}

func main()  {
	// Step 1: Unbuffered integer channel create kiya
	ch := make(chan int)

	// Step 2: Background Goroutine launch ki marks values (80, 75, 90) aur channel ke sath
	go calcTotal(80,75,90,ch) 

	// Step 3: Main Goroutine channel se calculation result aane ka wait karke store karegi
	// Note: 'total := <-ch' line BLOCK rehti hai jab tak background goroutine write nahi kar deti
	total := <- ch 
	// Step 4: Received total marks print kiya
	fmt.Println("Total:", total)


}




