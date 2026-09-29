package main

import "fmt"

// // 1. Struct Definition

type Student struct{
	Name string
	Marks float64

}

// 2. Normal Function

func printStudentInfo(s Student)  {
	fmt.Println("[Function Output] Name:",s.Name,"| Marks:", s.Marks)
}

// 3. Method (Attached to Student Struct via Receiver)
// Notice: (s Student) is the receiver
func (s Student) display() {
	fmt.Println("[Method Output] Name:",s.Name,"| Marks:",s.Marks)
}


func main()  {
	
	// Object Creation 
	student := Student{
		Name: "Arun",
		Marks: 85.9,
	}


    // ---------------------------------------------
   // Way 1: Calling Normal Function
   // ---------------------------------------------

	printStudentInfo(student)

	// ---------------------------------------------
   // Way 2: Calling Method
   // ---------------------------------------------

   student.display() // Method call Syntax: object.methodName() 





}


