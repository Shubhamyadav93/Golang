package main

import "fmt"

type Animal interface{
	Sound()
}

type Dog struct{}

func (d Dog) Sound() {
	fmt.Println("Dog says: Bark")
}

type Cat struct{}

func (c Cat) Sound() {
	fmt.Println("Cat says: Meow")
}

func makeSound(a Animal)  {
	a.Sound()
}


func main()  {
	makeSound(Dog{})
	makeSound(Cat{})
}

// Dog says: Bark
// Dog says: Meow

