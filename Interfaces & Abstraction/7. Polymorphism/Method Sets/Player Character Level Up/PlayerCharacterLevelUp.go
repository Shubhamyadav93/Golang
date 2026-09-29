package main

import "fmt"

// 1. Define the struct
type Player struct{
	Name string
	Level int 
}

// 2. Value Receiver: gets a copy, cannot change the original Player
func (p Player) show() {
	fmt.Printf("%s is at level %d\n",p.Name,p.Level)
}

// 3. Pointer Receiver: gets the address, CAN change the original Player
func (p *Player) Levelup() {
	p.Level++
}

func main()  {
	var playerName string
	// Read the player's name from input
	fmt.Scan(&playerName)

	// 4. Create a pointer to a Player
	// We use the '&' operator to get the memory address

	playerPts := &Player{
		Name: playerName,
		Level: 1,
	}

	// 5. Call the pointer receiver method
	// This modifies the original struct data
	playerPts.Levelup()
	// 6. Call the value receiver method
	// Pointers have access to both value and pointer methods!
	playerPts.show()

}



