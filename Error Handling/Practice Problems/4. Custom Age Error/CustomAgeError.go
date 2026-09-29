package main

import "fmt"

// 1. Custom Error Struct holding invalid age metadata

type AgeError struct{
	Age int 
}

// 2. Implementing built-in error interface 

func (e *AgeError) Error() string{
	return fmt.Sprintf("invalid age: %d",e.Age)
}


// Validation function returning custom error struct pointer 
func validationAge(age int) error {
	if age < 18 {
		// Instantiating custom struct error pointer 
		return &AgeError{Age: age}
	}
	return nil
}

func main()  {
	// Calling with invalid age = 15 
	err := validationAge(15)

	if err != nil {
		fmt.Println(err)
	}

}



