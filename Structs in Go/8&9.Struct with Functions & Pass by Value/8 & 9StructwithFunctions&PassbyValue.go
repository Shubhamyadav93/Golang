package main

import "fmt"

// Struct Definition

type Student struct{
	Name string
	Marks float64
}

// Function jo Struct ko Pass by Value ke roop me leta hai

func updateMarks(s Student)  {
	// s.Marks = 100
	s.Marks = 100 
}

func main()  {
	// Original Struct Object Creation 
	student := Student{
		Name: "Arun",
		Marks: 80,
	}
	fmt.Println("Before function Call (Original):",student.Marks)// Output: 80

	// Function call (Pass by Value - original ki copy banegi)
	updateMarks(student)

	// Function call ke baad original value check karna
	fmt.Println("After Fiunction Call (Original):",student.Marks)// 80 

}




