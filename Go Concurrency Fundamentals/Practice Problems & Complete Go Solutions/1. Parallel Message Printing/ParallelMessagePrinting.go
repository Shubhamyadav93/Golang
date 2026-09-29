package main

import (
	"fmt"
	"time"
)

// Har message ko 3 baar print karne wala function 
func printMsg(msg string)  {
	for i := 1; i<= 3; i++{
		fmt.Println(msg)
		time.Sleep(100 * time.Millisecond)// Har print ke baad 100ms pause 
	}
}

func main()  {
	// 3 alag-alag background Gorutines launch ho rahi hain 
	go printMsg("Hello")		// Gorutines 1     
	go printMsg("Welcome")		//Gorutines  2 
	go printMsg("Go Programming")		// Gorutines 3 
	
	// Main Goroutine ko 500ms tak roka taaki background tasks poore ho sakein 
	time.Sleep(500 * time.Millisecond)
}



