package main

import "fmt"

// 1. Contract / Interface

type Animal interface{
	Speak() string
}

// 2. Concrete Types (Structs)

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

// 3. Polymorphic Function
func makeSound(a Animal)  {
	fmt.Println(a.Speak())
}

func main()  {
	myDog := Dog{}
	myCat := Cat{}

   // Same function alag-alag types ke sath kaam kar raha hai
	makeSound(myDog)
	makeSound(myCat)

}





