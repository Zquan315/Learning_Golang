package channel_immediate

import (
	f "fmt"
	t "time"
)

func channel1(ch1 chan string) {
	arr := []string{"A1", "B1", "C1"}
	for _, v := range arr {
		ch1 <- v
		t.Sleep(1 * t.Second)
	}
}

func channel2(ch2 chan string) {
	arr := []string{"A2", "B2", "C2", "D2"}
	for _, v := range arr {
		ch2 <- v
		t.Sleep(2 * t.Second)
	}
}
func RunImmediateChannel() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go channel1(ch1)
	go channel2(ch2)

	for {
		flag := false
		select {
		case msg1 := <-ch1:
			f.Println("Received from channel 1:", msg1, "at", t.Now().Format("15:04:05"))
		case msg2 := <-ch2:
			f.Println("Received from channel 2:", msg2, "at", t.Now().Format("15:04:05"))
		case <-t.After(5 * t.Second):
			f.Println("No messages received within 5 seconds, exiting.")
			flag = true
		}

		if flag {
			break
		}
	}
	f.Println("⭐ Run Immediate Channel finished\n------------")
}
