package main

import "fmt"


func startApplication(config string)  {
	if config == ""{
		// Triggers an immediate fatal runtime panic due to missing critical setup 
		panic("configuration is missing")
	}
	fmt.Println("Application started")
}

func main()  {
	// Attempting startup with an empty config string ""
	startApplication("")
}


