package main

import "fmt"

func main()  {
	// Step 1: Do alag-alag String channels create kiye 
	ch1 := make(chan string)
	ch2 := make(chan string)


	// Step 2: Pehli Gorutine background me ch1 me data bhejti hai 

	go func ()  {
		ch1 <- "Hello"
	}()

	//  Step 3: Doosri Gorutine background me ch2 me data bhejti hai 
	go func ()  {
		ch2 <- "World"
	}()

	//  Step 4: select statement redy channel ka wait karti hai 

	select {
	case msg := <- ch1:
		fmt.Println(msg)
	case msg := <- ch2:
		fmt.Println(msg)
	}

}



