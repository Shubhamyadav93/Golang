package main

import "fmt"

func startApplication(config string)  {
	// 🛡️ Defer function acting as safety net
	defer func ()  {
		// recover() intercepts active panic and retrieves its message
		if r := recover(); r != nil{
			fmt.Println("Recovered from panic:",r)
		}

	}() 

	if config == "" {
		// Panic is triggered here 
		panic("configuration is missing")
	}
	fmt.Println("Application started")
}

func main()  {
	// Calling with empty configuration 
	startApplication("")
}


// Recovered from panic: configuration is missing


