package main

import "fmt"

type Notifier interface{
	Notify(message string) string
}

type Email struct{}

type SMS struct{}

func (e Email) Notify(message string) string {
	return "Sending Email:"+message
}
func (s SMS) Notify(message string) string {
	return "Sending SMS:" + message
}

func SendAlert(n Notifier,message string)  {
	fmt.Println(n.Notify(message))
}

func main()  {
	var notify string
	fmt.Scan(&notify)

	email := Email{}
	sms := SMS{}
	
	SendAlert(email,notify)
	SendAlert(sms,notify)
}



