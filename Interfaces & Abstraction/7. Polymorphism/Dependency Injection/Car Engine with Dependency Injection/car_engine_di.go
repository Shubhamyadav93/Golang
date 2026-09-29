package main

import "fmt"

// ==========================================
// 1. INTERFACE (Contract for Abstraction)
// ==========================================
type Engine interface {
	Start()
	FuelType() string
}

// ==========================================
// 2. CONCRETE STRUCTS (Implementations)
// ==========================================

// Struct 1: Gas Engine
type GasEngine struct {
	Cylinders int
}

// Value Receiver (Sirf read kar raha hai)
func (g GasEngine) FuelType() string {
	return "Petrol"
}

// Value Receiver (Polymorphic method 1)
func (g GasEngine) Start() {
	fmt.Printf("Vroom! %d-cylinder Gas engine started.\n", g.Cylinders)
}

// Struct 2: Electric Engine
type ElectricEngine struct {
	BatteryHealth int
}

// Value Receiver
func (e ElectricEngine) FuelType() string {
	return "Electric Battery"
}

// Pointer Receiver: Original BatteryHealth modify karne ke liye (Method Sets in action)
func (e *ElectricEngine) DegradeBattery(amount int) {
	e.BatteryHealth -= amount
}

// Value Receiver (Polymorphic method 2)
func (e ElectricEngine) Start() {
	fmt.Printf("Silent hum... Electric engine running at %d%% battery health.\n", e.BatteryHealth)
}

// ==========================================
// 3. DEPENDENCY INJECTION (Car Struct)
// ==========================================
type Car struct {
	Brand    string
	MyEngine Engine // Dependency injected via Interface
}

// Pointer Receiver method: Car brand modify karne ke liye
func (c *Car) Rebrand(newBrand string) {
	c.Brand = newBrand
}

// Value Receiver method: Car drive karne ke liye
func (c Car) Drive() {
	fmt.Printf("Driving %s (%s):\n", c.Brand, c.MyEngine.FuelType())
	c.MyEngine.Start() // Polymorphism: Jo bhi engine inject hoga, uska Start chalega
	fmt.Println("Car is moving smoothly!\n")
}

// ==========================================
// 4. MAIN WORKFLOW
// ==========================================
func main() {
	// Step A: Dependencies banayi
	v8 := GasEngine{Cylinders: 8}
	evMotor := &ElectricEngine{BatteryHealth: 100} // Pointer create kiya taaki degrade method chale

	// Step B: Battery modify ki (Pointer receiver demo)
	evMotor.DegradeBattery(10) // Ab 90% ho jayega

	// Step C: Dependency Injection (Engine ko Car ke andar inject kiya)
	car1 := Car{
		Brand:    "Ford Mustang",
		MyEngine: v8, // GasEngine injected
	}

	car2 := Car{
		Brand:    "Tesla Model S",
		MyEngine: evMotor, // *ElectricEngine injected
	}

	// Step D: Rebrand demo (Pointer receiver on Car)
	car1.Rebrand("Shelby GT")

	// Step E: Polymorphic execution
	car1.Drive()
	car2.Drive()
}

// Driving Shelby GT (Petrol):
// Vroom! 8-cylinder Gas engine started.
// Car is moving smoothly!

// Driving Tesla Model S (Electric Battery):
// Silent hum... Electric engine running at 90% battery health.
// Car is moving smoothly!


