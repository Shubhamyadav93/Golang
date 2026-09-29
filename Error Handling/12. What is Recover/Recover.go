package main

import "fmt"


func main()  {
	//1. Defer function setup (Safety Net Layer)

	defer func ()  {
		// recover() intercepts panic if active , otherwise returns nil 
		if r := recover(); r != nil{
			fmt.Println("Successfully Recovered from panic:",r)
		}
	}()

	fmt.Println("Aplication started successfully...")
	
	// 2. Triggering Panic 
	panic("unexpected error occured in system")
	
	//  // ⚠️ This line will be skipped because panic triggers immediate stack unwind
	fmt.Println("This line will never execute") 

}

// Aplication started successfully...
// Successfully Recovered from panic: unexpected error occured in system

