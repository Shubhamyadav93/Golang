package main

import "fmt"

// 1. Socket (Interface) banaya jo behaviour tay karta hai
type WallSocket interface{
	ConnectPlug()
}

// 2. Do alag types (Structs) banaye

type TV struct{}
type LaptopCharger struct{}

// 3. TV ne Plug lagane ka tarika define kiya (Interface implement ho gaya)
func (t TV) ConnectPlug() {
	fmt.Println("TV chalu ho gya aur screen on ho gayi!")
}
// 4. LaptopCharger ne bhi wahi same method define kiya
func (l LaptopCharger) ConnectPlug() {
	fmt.Println("Laptop charge hona shuru ho gaya!")
}

// 5. Yeh function kisi bhi WallSocket ke sath kaam kar sakta hai
func PowerUp(device WallSocket)  {
	device.ConnectPlug()// Yeh bina jaane chalega ki device TV hai ya Charger
}

func main()  {
	myTV := TV{}
	myCharger := LaptopCharger{}

	PowerUp(myTV) // Output: TV chalu ho gaya...
	PowerUp(myCharger) // Output: Laptop charge hona...

}

