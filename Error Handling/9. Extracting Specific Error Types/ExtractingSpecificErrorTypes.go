package main

import (
	"errors"
	"fmt"
)

// 1. Custom Error Struct (Stores extra metadata like Code)

type SystemError struct{
	Code int 
	Message string
}

// Implementing error interface for SystemError 

func (e *SystemError) Error() string {
	return fmt.Sprintf("Enter Code %d: %s", e.Code,e.Message)
}

// Function returning a wrapped custom error 

func ExecuteTask() error {
	sysErr := &SystemError{
		Code: 500,
		Message: "Internal Server Failure",
	}
	// Wrapping custom error 
	return fmt.Errorf("task failed: %w",sysErr)
}

func main()  {
	err := ExecuteTask()

	// Target pointer variable created to hold the extracted custom struct 
	var sysErr *SystemError 
	// error.AS check if 'err' contains '*SystemError' and extracts it into 'sysErr'

	if errors.As(err,&sysErr){
		fmt.Println("Custom Error Found!")
		fmt.Println("Extracted Status Code:",sysErr.Code) // Accessing struct feild!
	}else{
		fmt.Println("Not a SystemError type")
	}
}


