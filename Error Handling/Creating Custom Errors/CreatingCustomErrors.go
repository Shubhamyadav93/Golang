package main

import (
	"errors"
	"fmt"
)

// Custom validation function error

func CheckAge(age int) error {
	if age < 18 {
		// Returing custom error message for domain rule violation 
		return errors.New("age must be at least 18")
	}
	// Success case: No error occurred 
	return nil 
}

func main()  {
	// Calling validation with invalid input (16 < 18)
	err := CheckAge(16)
	// Guard Clause checking if custom error was returned 
	if err != nil {
		fmt.Println("Validation Error:",err)
	}
}



// Validation Error: age must be at least 18




