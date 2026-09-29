package student

import "fmt"

type Student struct{
	Name string
	RollNo int 
}

func (s Student) Display() {
	fmt.Printf("Student: %s (Roll: %d)",s.Name,s.RollNo)
}


