package main

import "fmt"

// calculate function ek channel argument leta hai

func calculate(ch chan int)  {
	result := 10 + 10 
	ch <- result // Step 2: Calculation ke baad result channel e bhej diya 
}

func main()  {
	// Step 1 : Integer type ka unbuffered channel create ki 
	ch := make(chan int)

	// Step 2: Background me calculate gorutine start ki 

	go calculate(ch) 

	// Step 3: Channel se data receive kiya aur 'result' variable me rakha 

	result := <- ch 
	fmt.Println("Result:", result)
}

