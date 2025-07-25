package miniproject

import (
	f "fmt"
	"math/rand"
	"sync"
	t "time"
)

func calcx10(arr []*int, ch1 chan int, wg *sync.WaitGroup, mu *sync.Mutex) {
	for i, v := range arr {
		t.Sleep(t.Duration(rand.Intn(1000)+500) * t.Millisecond)
		if *v < 10 {
			n := *v
			mu.Lock()
			*v *= 10
			*arr[i] = *v
			ch1 <- *v
			f.Println("Calcx10 with:", n, "at", t.Now().Format("15:04:05"))
			mu.Unlock()
		}

	}
	wg.Done()
}
func calcx100(arr []*int, ch2 chan int, wg *sync.WaitGroup, mu *sync.Mutex) {
	for i, v := range arr {
		t.Sleep(t.Duration(rand.Intn(1000)+500) * t.Millisecond)
		if *v < 10 {
			n := *v
			mu.Lock()
			*v *= 100
			*arr[i] = *v
			ch2 <- *v
			f.Println("Calcx100 with:", n, "at", t.Now().Format("15:04:05"))
			mu.Unlock()
		}
	}
	wg.Done()
}
func RunMiniProject() {
	rand.Seed(t.Now().UnixNano())
	ch1 := make(chan int)
	ch2 := make(chan int)
	mu := new(sync.Mutex)
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(2)
	arr := make([]*int, 9)
	for i := 0; i < 9; i++ {
		arr[i] = new(int)
		*arr[i] = i + 1
	}
	go calcx10(arr, ch1, wg, mu)
	go calcx100(arr, ch2, wg, mu)

	for {
		flag := false
		select {
		case v := <-ch1:
			f.Println("Value from calcx10 - channel 1:", v)
		case v := <-ch2:
			f.Println("Value from calcx100 - channel 2:", v)
		case <-t.After(5 * t.Second):
			f.Println("No more values received, exiting...")
			flag = true
		}
		if flag {
			break
		}
	}
	wg.Wait()
	f.Println("Final array values:")
	for _, v := range arr {
		f.Print(*v, " ")
	}
	f.Println("\n⭐ Run Mini Project finished\n------------")
}
