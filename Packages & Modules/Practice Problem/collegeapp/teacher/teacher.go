package teacher

import "fmt"

type Teacher struct{
	Name string
	Subject string
}

func (t Teacher) Display() {
	fmt.Printf("Teacher: %s (Subject: %s)",t.Name,t.Subject)
}
