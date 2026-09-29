package main

import "fmt"

type Storage interface{
	Save(data string)
}

type CloudStorage struct{}

func (c CloudStorage) Save(data string) {
	fmt.Println("Saving data to clud Storage")
	fmt.Println("Data:",data)
}


func saveData(s Storage,data string)  {
	s.Save(data)
}

func main()  {
	saveData(CloudStorage{},"Student Registration Data")
}


