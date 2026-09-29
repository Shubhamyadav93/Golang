package main

import "goproject/students"


func main()  {
	s:= students.Student{
		Name: "Arun",
		Age: 20,
		Department: "CSE",
	}
	students.DisplayStudent(s)
}


// Student Details
// Name: Arun
// Age: 20
// Department: CSE 

