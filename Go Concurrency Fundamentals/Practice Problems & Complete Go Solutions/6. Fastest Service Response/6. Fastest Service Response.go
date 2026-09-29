package main

import (
	"fmt"
	"time"
)

func main()  {
	// Step 1: Do channels create kiye alag-alag services ke liye
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Step 2: Service 1 (Slow Service - 3 Seconds)
	go func ()  {
		time.Sleep(3 * time.Second)
		ch1 <- "Response from Service 1"
	}()

	// Step 3: Service 2 (Fast Service - 1 Second)
	go func ()  {
		time.Sleep(1 * time.Second)
		ch2 <- "Response from Service 2"
	}()

	// Step 4: select block jo fast response ka wait karega

	select {
	case res := <- ch1:
		fmt.Println(res)

	case res := <- ch2:
		fmt.Println(res)
	}

}


