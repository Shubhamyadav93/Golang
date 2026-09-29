package main

import (
	"fmt"
	"time"
)


func main()  {
	// Step 1: Unbuffered string channel create kiya

	ch := make(chan string)

	// Step 2: Background Goroutine launch ki jo 3 seconds baad data bheje gi
	go func ()  {
		time.Sleep(3 * time.Second) // 3 seconds delay simulate kiya
		ch <- "Data received successfully" // Channel me msg send kiya
	}()


	// Step 3: Select block se channel response aur timeout ke beech race hoti hai
	select{
	case msg := <- ch:
		 // Agar background goroutine 2 seconds se PEHLE data bhej deti
		 fmt.Println(msg)

	case <- time.After(2 * time.Second):
		// Agar 2 seconds tak channel 'ch' me koi data NAHI aaya
		fmt.Println("Request timed out")
	}

}



