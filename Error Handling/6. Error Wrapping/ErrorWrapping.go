package main

import (
	"errors"
	"fmt"
)

// Lower-level function (Database/File/I/O layer)

func ReadFile() error {
	return errors.New("file not found")
}

// Mid-level function (Business logic layer) that wraps the error 

func LoadProfile() error {
	err := ReadFile()
	if err != nil{
		// %w verb wraps the original error 'err' inside the new contextual error 
		return fmt.Errorf("failed to loas profile: %w",err)
	}
	return nil

}

func main()  {
	// Top-level caller 
	err := LoadProfile()
	if err != nil{
		fmt.Println("Handled Error:",err)
	}
}



// Handled Error: failed to loas profile: file not found


