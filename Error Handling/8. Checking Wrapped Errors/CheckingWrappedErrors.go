package main

import (
	"errors"
	"fmt"
)

// Global Sentinel Error (Standard predefined error)
var ErrNotFound = errors.New("record not found")

// Database layer function jo error ko wrap karti hai 

func FindUser() error {
	// ErrNotFound ko wrap karke naya message banaya 
	return fmt.Errorf("database error: %w",ErrNotFound)
}


func main()  {
	err := FindUser()
  // ❌ GALAT APPROACH: String matching fail ho jayegi
   // if err.Error() == "record not found" { ... } // FALSE! Kyunki text hai "database error: record not found"

  // 🟢 SAHI APPROACH: Go ka errors.Is() internal chain me jakar match karega

  if errors.Is(err,ErrNotFound){
	fmt.Println("User not found (Matched via errors.Is!)")
  }

}



