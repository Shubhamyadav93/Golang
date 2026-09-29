package main

import "fmt"

// Step 1: Define the Interface
type Weapon interface {
	Use() string
}

// Step 2: Create the Dependencies
type Sword struct{}

type Bow struct{}


// Step 3: Implement the Method Sets
func (s Sword) Use() string {
	return "Swung a sword!"
}

func (b Bow) Use() string {
	return "Fired an arrow!"
}

// Step 4: Create the Main Struct
type Hero struct {
	EquippedWeapon Weapon
}

// Step 5: Use the Injected Dependency
func (h Hero) Attack() {
	fmt.Println("Hero attacked:",h.EquippedWeapon.Use())
}

func main() {
	var input string
	fmt.Scan(&input)
	var weapon Weapon
	
	if input == "bow"{
	        weapon = Bow{}
	}else{
	        weapon = Sword{}
	}
	
	hero := Hero{
	        EquippedWeapon:weapon,
	}
	hero.Attack()
}

// Hero attacked: Swung a sword!
