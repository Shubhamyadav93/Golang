package main

import "fmt"

type Device interface{
	TuneOn() string
}

type Light struct{}

type Fan struct{}

func (l Light) TuneOn() string {
	return "Light is illuminating the room!"
}

func (f Fan) TuneOn() string {
	return "Fan is colling the room!"
}

func ActivateDevice(d Device)  {
	fmt.Println(d.TuneOn())
}

func main()  {
	myLight := Light{}
	myFan := Fan{}

	ActivateDevice(myLight)
	ActivateDevice(myFan)
}
