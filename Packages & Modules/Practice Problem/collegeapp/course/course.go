package course

import "fmt"

type Course struct{
	Code string
	Title string
}

func (c Course) Display() {
	fmt.Printf("Course: [%s] %s",c.Code,c.Title)
}
