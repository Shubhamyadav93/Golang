package main

import (
	"errors"
	"fmt"
)

// Lower-level layer: Creates original root error
func ReadFile() error {
	return errors.New("file not found")
}

// Mid-level layer: Wraps original error using %w 

func LoadProfile() error {
	err := ReadFile()
	if err != nil{
		return fmt.Errorf("failed to load profile: %w",err)
	}
	return nil
}


func main()  {
	// Top-level caller receive wrapped error 
	wrappedErr := LoadProfile()
	fmt.Println("Wrapped Error :",wrappedErr)

	// UNWRAPPING: Extracting the original root error 

	originalErr := errors.Unwrap(wrappedErr)
	fmt.Println("Unwrapped Error:",originalErr)

}


// Unwrapped Error: file not found 


