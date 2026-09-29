package main

import (
	"fmt"
	"time"
)

// Worker Function: Continuous job listening aur shutdown control ke liye
func worker(jobCh chan int,shutdownCh chan bool)  {
	for{
		select {
		case job := <- jobCh:
			// Continuous state: Jobs process karta hai
			fmt.Printf("Processing Job: %d\n",job)

		case <- shutdownCh:
			// Exit state: Signal milne par clean shutdown karta hai
			fmt.Println("Worker shutting down")
			return // Infinite loop se bahar nikal kar Goroutine exit karta hai 

		}
	}
}

func main()  {
	// Step 1: Data aur Control channels initialize kiye
	jobCh := make(chan int)
	shutdownCh := make(chan bool)

	// Step 2: Background worker Goroutine launch ki
	go worker(jobCh,shutdownCh)

	// Step 3: Worker ko 2 jobs bheje
	jobCh <- 101 
	jobCh <- 102 

	time.Sleep(100 * time.Microsecond)// Jobs execution allow karne ke liye wait delay

	// Step 4: Worker ko Shutdown Signal bheja
	shutdownCh <- true 

	time.Sleep(100 * time.Microsecond)// Worker execution complete hone ka buffer time

}



