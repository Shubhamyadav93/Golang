package main

import "fmt"

type Shape interface{
	Name() string
}

type Circle struct{}
type Square struct{}

func (c Circle) Name() string {
	return "Circle"
}


func (s Square) Name() string {
	return "Square"
}


func main()  {
	shapes := []Shape{
		Circle{},
		Square{},
	}

	// Step 5: Print shape names
	fmt.Println("Shape 1:", shapes[0].Name())
	fmt.Println("Shape 2:",shapes[1].Name())

}

// Shape 1: Circle
// Shape 2: Square

