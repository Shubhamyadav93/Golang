package main

import "fmt"

// 1. INTERFACE (Common Rule)

type Animal interface{
	Sound()
}

// 2. STRUCTS (Alag-Alag Animals)

type Dog struct{}
type Cat struct{}
type Cow struct{} 

// 3. IMPLEMENTATION (Har Animal ka Sound alag behavior deta hai)

func (d Dog) Sound() {
	fmt.Println("Dog says: Bark! 🐶")
}

func (c Cat) Sound() {
	fmt.Println("Cat says: Meow! 🐱")

}

func (cw Cow) Sound() {
	fmt.Println("Cow says: Moo! 🐮")

}

// 4. POLYMORPHIC FUNCTION
// Yeh function kisi bhi Animal ko sound nikalne ki instruction deta hai
func MakeAnimalSound(a Animal)  {
	a.Sound() // Instruction SAME hai (Sound), par behavior alag hai!
}


func main()  {
	// Same function ko alag-alag Struct objects de rahe hain:

	MakeAnimalSound(Dog{}) // Output :- Dog says: Bark! 🐶
	MakeAnimalSound(Cat{}) //Output :- // Cat says: Meow! 🐱
	MakeAnimalSound(Cat{}) // Output :- // Cat says: Meow! 🐱
}

// Dog says: Bark! 🐶
// Cat says: Meow! 🐱
// Cat says: Meow! 🐱

