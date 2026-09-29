package main

import "fmt"

type Student struct {
	Name       string
	Age        int
	Department string
	Marks      float64
}

func main() {
	s := Student{
		Name:       "Arun",
		Age:        20,
		Department: "CSE",
		Marks:      85.5,
	}
	fmt.Println("Student Details:")
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
	fmt.Println("Department:",s.Department)
	fmt.Println("Marks:",s.Marks)

}

// Student Details:
// Name: Arun
// Age: 20
// Department: CSE
// Marks: 85.5
