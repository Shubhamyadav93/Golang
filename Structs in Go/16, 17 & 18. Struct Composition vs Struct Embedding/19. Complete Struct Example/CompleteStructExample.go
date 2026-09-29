package main

import "fmt"


type Address struct{
	City string
	State string
}

type Student struct{
	Name string
	Age int 
	Marks float64
	Address Address
}

// Value Receiver Method (display)
func (s Student)display() {
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
	fmt.Println("Marks:",s.Marks)
	fmt.Println("City:",s.Address.City)
}

// Pointer Receiver Method (updateMarks)

func (s *Student) updateMarks(marks float64){
	s.Marks = marks 
}

func main()  {

	// Object Create
	student:= Student{
		Name:"Arun",
		Age:20,
		Marks:85,
		Address:Address{
			City:"Chennai",
			State:"Tamil Nadu",
		},
	}

	student.display()
	student.updateMarks(95)
	fmt.Println("Update Marks:",student.Marks)

}

// Name: Arun
// Age: 20
// Marks: 85
// City: Chennai
// Update Marks: 95



