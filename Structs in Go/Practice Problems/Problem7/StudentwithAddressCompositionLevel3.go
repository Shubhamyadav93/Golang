package main

import "fmt"

type Address struct{
	City string
	State string
	Pincode string
}

type Student struct{
	Name string
	Age int 
	Department string
	Addr Address 
}

func ( s Student) displayDetails() {
	fmt.Println("Student Details:")
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
	fmt.Println("Department:",s.Department)
	fmt.Println("Address:")
	fmt.Println("City:",s.Addr.City)
	fmt.Println("State:",s.Addr.State)
	fmt.Println("Pincode:",s.Addr.Pincode)
}

func main()  {
	student := Student{
		Name: "Divya",
		Age: 20,
		Department: "CSE",
		Addr: Address{
			City: "Combatore",
			State: "Tamil Nadu",
			Pincode: "641001",
		},
	}

	student.displayDetails()
}

// Student Details:
// Name: Divya
// Age: 20
// Department: CSE
// Address:
// City: Combatore
// State: Tamil Nadu
// Pincode: 641001


