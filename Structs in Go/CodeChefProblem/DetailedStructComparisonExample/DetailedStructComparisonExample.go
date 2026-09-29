package main

import "fmt"

type Student struct{
	Name string
	Marks int 
}

// 1. VALUE RECEIVER (Read-only behavior)
// Yeh method data ki copy par kaam karta hai

func (s Student) UpdateMarksValue(newMarks int) {
	s.Marks = newMarks // Yeh modification sirf is method ke andar tak limited rahega
}


// 2. POINTER RECEIVER (Mutable behavior)
// *Student batata hai ki yeh original memory reference ko point kar raha hai
func (s *Student) UpdateMarksPointer(newMarks int) {
	s.Marks = newMarks // Yeh original variable ko update kar dega
}

// 3. VALUE RECEIVER (Sirf display/print karne ke liye best)
func (s Student) Display() {
	fmt.Printf("Name:%s | Marks:%d\n",s.Name,s.Marks)
}

func main()  {
	stud := Student{
		Name: "Rohan",
		Marks: 50,
	}
   // Step 1: Value receiver test

   stud.UpdateMarksValue(80)
   stud.Display()
   // Output: Marks 50 hi rahenge, kyunki copy modify hui thi


      // Step 2: Pointer receiver test

	stud.UpdateMarksPointer(95)
	stud.Display()
   // Output: Marks 95 ho jayenge, kyunki original address par value change hui

}

// Name:Rohan | Marks:50
// Name:Rohan | Marks:95

