package main

import (
	"fmt"
	"time"
)

// Helper function jo name aur loop count print karega
func task(name string)  {
	for i := 1; i <= 3; i ++{
		fmt.Println(name,i)
		time.Sleep(500 * time.Millisecond) // 0.5 second pause
	}
}

func main()  {
	// Task A aur Task B dono ko Goroutines me start kar rahe hain
	go task("Task A")
	go task("Task B")
	time.Sleep(2 * time.Second) // Main Goroutine ko 2 seconds ke liye roka taaki background tasks poore ho sakein

}


