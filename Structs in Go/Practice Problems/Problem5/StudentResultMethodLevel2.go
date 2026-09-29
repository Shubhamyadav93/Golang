package main

import "fmt"

type Student struct {
	Name  string
	Marks float64
}

func (s Student) displayResult() {
	fmt.Println("Student:",s.Name)
	if s.Marks >= 40 {
		fmt.Println("Result: Pass")
	}else{
		fmt.Println("Result: Fail")
	}
}


func main()  {
	student := Student{
		Name: "Arun",
		Marks: 75,
	} 

	student.displayResult()
}

// Student: Arun
// Result: Pass 



