package goroutine_channel

import (
	f "fmt"
	"time"
)

func stringValue(ch chan string) {
	s := [3]string{"Hello", "World", "Golang"}
	for _, value := range s {
		ch <- value
		time.Sleep(500 * time.Millisecond)
	}
}

func speValue(ch chan string) {
	s := [3]string{"&", "*", "$"}
	for _, value := range s {
		ch <- value
		time.Sleep(500 * time.Millisecond)
	}
}

func RunBasicChannel() {
	ch := make(chan string)
	go func() {
		stringValue(ch)
		speValue(ch)
		close(ch)
	}()
	for val := range ch {
		f.Println(val)
	}

	f.Println("⭐ Run Basic Channel finished\n------------")
}
