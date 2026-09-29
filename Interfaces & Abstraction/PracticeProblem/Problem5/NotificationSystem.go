package main

import "fmt"

type Notification interface{
	Send(message string)
}


type Email struct{
	EmailID string
}

func (e Email) Send(msg string) {
	fmt.Printf("Sending Email to %s | Message: %s\n", e.EmailID, msg)
}


type SMS struct{
	Phone string
}


func (s SMS) Send(msg string) {
	fmt.Printf("Sending SMS to %s | Message: %s\n", s.Phone, msg)
}


type WhatsApp struct{
	WhatsAppmsg string
} 


func (w WhatsApp) Send(msg string) {
	fmt.Printf("Sending WhatsApp to %s | Message: %s\n", w.WhatsAppmsg, msg)
}


func sendNotification(n Notification,message string)  {
	n.Send(message)
}

func main()  {
	em := Email{
		EmailID: "student@example.com",
	}
	fmt.Printf("----------Message for Email-------------- ")
	sendNotification(em,"Welcome to the course")

	sm := SMS{
		Phone: "4784274947",
	}
	fmt.Printf("-----------Message for SMS----------- ")
	sendNotification(sm,"Welcome to the newCource")
	
	wh := WhatsApp{
		WhatsAppmsg: "+18357429765",
	}
	fmt.Printf("--------------Message for WhatsApp----------")
	sendNotification(wh,"Welcome to Backend Course")

}





// ----------Message for Email-------------- Sending Email to student@example.com | Message: Welcome to the course
// -----------Message for SMS----------- Sending SMS to 4784274947 | Message: Welcome to the newCource
// --------------Message for WhatsApp----------Sending WhatsApp to +18357429765 | Message: Welcome to Backend Course


