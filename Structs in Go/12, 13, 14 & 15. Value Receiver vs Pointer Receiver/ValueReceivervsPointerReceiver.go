package main

import "fmt"

type Student struct{
	Name string
	Marks float64
}

// // Value Receiver Definition

func (s Student) updateMarksValue(marks float64) {
	s.Marks = marks // Sirf COPY modify hoti hai

}

// // Pointer Receiver Definition

func (s *Student) updateMarksPointer(marks float64){
	s.Marks = marks
}


func main()  {
	student := Student{
		Name: "Arun",
		Marks: 85,
	}

	student.updateMarksValue(90)
	fmt.Println(student.Marks) // Output: 80 (Original BADLA NAHI)

	student.updateMarksPointer(90)
	fmt.Println(student.Marks)// Output: 95 (Original UPDATE ho gaya!)

}


