package main

import (
	"errors"
	"fmt"
)

// Lower-level data layer function: Returns root cause error
func findStudent(id int) error {
	return errors.New("student not found")
}

// Service layer function: Wraps lower-level error with contextual info 

func getStudentDetails(id int) error {
	err := findStudent(id)
	
	if err != nil{
		// %w preserve the underlying 'student not found' error 
		return fmt.Errorf("failed to find student: %w",err)
	}
	return nil 
}

func main()  {
	//Attempting to fetch details for student ID 101 
	err := getStudentDetails(101)
	if err != nil{
		fmt.Println("Handled Error:",err)
	}
}


