package main

import "goproject1/students"

func main()  {
	s := students.Students{
		Name: "Rahul",
		Age: 20,
	}
	students.Display(s)
}
