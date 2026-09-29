package main

import "fmt"

// Step 1: Define the Interface
type Notifier interface{
	Send(message string)
}

// Step 2 & 3: Create EmailNotifier and its method set
type EmailNotifier struct{}
func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending Email:",message)
}

// Step 2 & 3: Create SMSNotifier and its method set
type SMSNotifier struct{}
func (s SMSNotifier) Send(message string) {
	fmt.Println("Sending SMS:",message)
}

// Step 4: Create the Main Struct with the interface as a dependency
type WeatherApp struct{
	AlertSystem Notifier
}

// Step 5: Use the dependency
func (w WeatherApp) Alert(message string) {
	w.AlertSystem.Send(message)
}

func main()  {
	// Read the weather condition from input
	var wetherCondition string
	fmt.Scan(&wetherCondition)

	// Create our dependencies
	emailNotif := EmailNotifier{}
	smsNotif := SMSNotifier{}

	// Create our apps and INJECT the dependencies
	app1 := WeatherApp{
		AlertSystem: emailNotif,
	}
	app2 := WeatherApp{
		AlertSystem: smsNotif,
	}

		// Trigger the alerts
	app1.Alert(wetherCondition)
	app2.Alert(wetherCondition)

}



// Blizzard
// Sending Email: Blizzard
// Sending SMS: Blizzard

