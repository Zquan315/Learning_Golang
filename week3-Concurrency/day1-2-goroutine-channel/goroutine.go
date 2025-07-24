package goroutine_channel

import (
	f "fmt"
	"time"
)

func printNumeric() {
	for i := 0; i < 5; i++ {
		f.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printString() {
	for i := 0; i < 5; i++ {
		f.Println("Quan", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printSpecial() {
	s := [5]string{"*", "%", "#", "@", "$"}
	for _, value := range s {
		f.Println(value)
		time.Sleep(500 * time.Millisecond)
	}
}
func RunGoRoutine() {
	go printNumeric()
	go printString()
	go printSpecial()
	time.Sleep(10 * time.Second)
	f.Println("⭐ Run GoRoutine finished\n------------")
}
