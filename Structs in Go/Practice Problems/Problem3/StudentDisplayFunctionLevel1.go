package main

import "fmt"

type Student struct{
	Name string
	Age int 
	Marks float64
}


func displayStudent(s Student)  {
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
	fmt.Println("Marks:",s.Marks)
}

func main()  {
	student:= Student{
		Name:"Priya",
		Age: 21,
		Marks: 92,
	}
	displayStudent(student)
} 

// Name: Priya
// Age: 21
// Marks: 92 

