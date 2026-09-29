package main

import "fmt"

type Device interface{
	On()
	Off() 
}


type TV struct{}

func (s TV) On() {
	fmt.Println("TV is On")
}

func (s TV) Off() {
	fmt.Println("TV is Off")
}

type AC struct{}

func (s AC) On()  {
	fmt.Println("AC is ON")
}

func (s AC) Off() {
	fmt.Println("AC is Off")
}

func main()  {
	var d Device // interface variable
	d = TV{} // TV implements Device
	
	d.On()
	d.Off()

	var d1 Device
	d1 = AC{}

	d1.On()
	d1.Off()


}

// TV is On
// TV is Off
// AC is ON
// AC is Off

