package main

import (
	"errors"
	"fmt"
)

// Domain validation function single error type
func validateMarks(mark int) error {
	// Boundary check: Mark must be in range [0,100]
	if mark < 0 || mark > 100 {
		return errors.New("marks must be between 0 and 100")
	}
	// Pass case 
	return nil
}

func main()  {
	// Attempting validation with invalid mark (105)
	err := validateMarks(105)

	// Signal Return Guard Check 
	if err != nil {
		fmt.Println("Validation Error:",err)
	}else{
		fmt.Println("Valid marks")
	}
}

// Validation Error: marks must be between 0 and 100





