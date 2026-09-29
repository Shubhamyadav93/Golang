package main

import "fmt"

type Shape interface{
	Area() float64
}

type Rectangle struct{
	Length, Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth 
}


type Circle struct{
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}


func calculateArea(s Shape)  {
	fmt.Printf("Area: %.2f\n", s.Area()) 
}

func main()  {
	r := Rectangle{
		Length: 10,
		Breadth: 5,
	}
	c := Circle{
		Radius: 7,
	}
	fmt.Println("Rectangle")
	calculateArea(r)

	fmt.Println("Circle ")
	calculateArea(c)
}

// Rectangle 
// Area: 50.00
// Circle 
// Area: 153.94%  


