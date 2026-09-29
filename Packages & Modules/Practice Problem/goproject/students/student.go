package students

import "fmt"

type Student struct{
	Name string
	Age int 
	Department string
}

func DisplayStudent(s Student)  {
	fmt.Println("Student Details")
	fmt.Println("Name:",s.Name)
	fmt.Println("Age:",s.Age)
	fmt.Println("Department:",s.Department)
}


