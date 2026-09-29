package main

import "fmt"


func main()  {
	// Declaration + Initialization ek sath (Bina type name ke) 
	user := struct{
		UserName string
		IsActive bool
	}{
		UserName: "arun_dev",
		IsActive: true,
	}

	// Accessing fields 
	fmt.Println("Username:",user.UserName)
	fmt.Println("Is Active:",user.IsActive)

}


