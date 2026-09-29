package main

import (
	"fmt"
	"runtime/debug"
	"strings"
)

type Notification interface{
	Format()
	Channel()
}

type Email struct{
	to string
	subject string
}

func (e Email) Format() string {
	return "Email to " + e.to + ": " + e.subject
}

func (e Email) Channel() string {
	return "Email"
}

type SMS struct{
	phone string
	text string
}

func (e SMS) Format() string {
	return "SMS to " + e.phone + ": " + e.text
}

func (e SMS) Channel()string {
	return "SMS"
}

type Push struct{
	device string
	text string
}

func (p Push) Format()string {
	return "Push to "+p.device+": " + p.text 
}

func (p Push) Channel() string {
	return "Push"
}

func main()  {
	var n int 
	fmt.Scan(&n)

	notifications := make([]Notification,0,n)

	for i:=0;i<n;i++{
		var typ string
		fmt.Scan(&typ)

		switch typ {
		case "E":
			var to,subject string
			fmt.Scan(&to,&subject)
			notifications = append(notifications,Email{
				to:to,
				subject: subject,
			}),

		case "S":
			var phone,text string
			fmt.Scan(&phone,&text)
			notifications = append(notifications,SMS{
				phone: phone,
				text:text,
			})

		case "P":
			var device ,text string
			fmt.Scan(&debug,&text)
			notifications=append(notifications,Push{
				device: device,
				text: text,
			})
		}
	}

	emailCount := 0 
	smsCount := 0
	pushCount := 0 

	for _,notification := range notifications{
		fmt.Println(notification.Format())

		switch notification.Channel(){
		case "Email":
			emailCount ++ 
		
		case "SMS":
			smsCount ++
		
		case "Push":
			pushCount ++
		}
	}
	fmt.Printf("Email: %d\n",emailCount)
	fmt.Printf("SMS:  %d\n",smsCount)
	fmt.Printf("Push: %d\n",pushCount)

}