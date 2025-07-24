package sync

import (
	f "fmt"
	"sync"
)

func func1(wg *sync.WaitGroup) {
	defer wg.Done()
	f.Println("Function 1 is running")
}
func func2(wg *sync.WaitGroup) {
	defer wg.Done()
	f.Println("Function 2 is running")
}
func func3(wg *sync.WaitGroup) {
	defer wg.Done()
	f.Println("Function 3 is running")
}
func usingWaitGroup() {
	defer f.Println("All functions completed with WaitGroup")
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	wg.Add(3)
	go func1(wg)
	go func2(wg)
	go func3(wg)
	wg.Wait()

}

var count int = 0

func increaseCountby2(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	count += 2
	mu.Unlock()
	f.Println("Count increased by 2, current count:", count)
}

func increaseCountby1(wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()
	mu.Lock()
	count += 1
	mu.Unlock()
	f.Println("Count increased by 1, current count:", count)
}

func usingMutex() {
	defer f.Println("All functions completed with Mutex")
	var wg *sync.WaitGroup = new(sync.WaitGroup)
	var mu *sync.Mutex = new(sync.Mutex)
	wg.Add(2)
	go increaseCountby1(wg, mu)
	go increaseCountby2(wg, mu)
	wg.Wait()
	f.Print("Current count: ", count, "\n")
}
func RunSync() {

	//usingWaitGroup()
	usingMutex()
	f.Println("⭐ Running Sync successfully!\n------------")
}
