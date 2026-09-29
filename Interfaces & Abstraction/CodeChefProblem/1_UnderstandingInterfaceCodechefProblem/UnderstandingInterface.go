package main

import (
	"fmt"
	"math"
)

// Step 1: Create the Interface (// 1. Interface (Niyam wahi purana hai))
type Shape interface{
	Area() float64
	Perimeter() float64 // Circle ke case mein perimeter ko Circumference bolte hain
}


// Step 2: Create a Struct (// 2. Square Struct aur uske methods)
type Square struct{
	side float64
}

// Step 3: Implement the Interface Methods

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}


// ==========================================
// NEW PART: 3. Circle Struct banaya
// ==========================================

type Circle struct{
	radius float64
}

// Circle ke liye Area ka formula: π * r * r
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle ke liye Perimeter (Circumference) ka formula: 2 * π * r
func (c Circle) Perimeter() float64  {
	return 2 * math.Pi * c.radius
}
func main()  {
		// Dono shapes ke real objects banaye
	sq := Square{
		side: 5,
	}
	ci := Circle{
		radius: 3,
	}

		// Ek akele 'Shape' type ka khali box variable banaya
	var sh Shape

		// --- PEHLA ROOP (SQUARE) ---
	sh = sq // Box ke andar Square daal diya
	fmt.Println("--- Square details using interface ---")
	fmt.Println("Area:", sh.Area())// Square ka Area chalega
	fmt.Println("Perimeter:",sh.Perimeter()) // Square ka Perimeter chalega

		// --- DOOSRA ROOP (CIRCLE) ---
	sh = ci 
	fmt.Println("\n--- Circle details using interface ---")
	fmt.Printf("Area: %.2f\n", sh.Area())    // Ab automatic Circle ka Area chalega!	
	fmt.Printf("Perimeter: %.2f\n", sh.Perimeter()) // Ab automatic Circle ka Perimeter chalega!


}


// --- Square details using interface ---
// Area: 25
// Perimeter: 20

// --- Circle details using interface ---
// Area: 28.27
// Perimeter: 18.85



// Main Function Mein Asli Khel Kya Hua?
// var sh Shape: Humne ek elastic box banaya jo kisi bhi Shape ko hold kar sakta hai.
// sh = sq karne ke baad jab humne sh.Area() chalaya, toh Go ne piche dekha ki box mein Square hai, toh usne square ka formula chalaya.
// sh = ci: Humne purana data nikal kar usi sh variable mein Circle daal diya. Go ne bina kisi error ke ise accept kar liya kyunki Circle ne bhi interface ke dono niyam (Area aur Perimeter) poore kiye the.
// Ab jab humne wahi same line sh.Area() dobara chalayi, toh is baar output badal gaya! Is baar Go ne pehchan liya ki box mein Circle baitha hai, toh usne circle ka formula chala diya.
// Isee ko Polymorphism kehte hain—variable ek hi hai (sh), line bhi wahi hai (sh.Area()), par dono baar kaam alag-alag ho raha hai uske andar baithe object ke hisaab se.


