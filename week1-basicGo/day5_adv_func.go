package main

import (
	"fmt"
	"os"
)

func learn_closure() func() int {
	count := 1
	return func() int {
		count *= 2
		return count
	}
}
func learn_closure_main() {
	init := learn_closure()
	for i := 1; i <= 3; i++ {
		fmt.Printf("Number %d: %d\n", i, init())
	}
}

func RunReadFile(filename string) {
	defer fmt.Print("Finished reading file.\n")
	read, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File content:", string(read))
}

func RunPAnic() {

	var a, b int
	fmt.Println("Enter two integers to divide (a/b):")
	fmt.Scan(&a, &b)
	defer recoverFromPanic(a, b)
	if b == 0 {
		panic(fmt.Sprintf("Division by zero: %d / %d", a, b))
	}
	result := a / b
	fmt.Printf("Result of %d / %d = %d\n", a, b, result)
}

func recoverFromPanic(a, b int) {
	if r := recover(); r != nil {
		fmt.Printf("Panic occurred: cannot divide %d by %d\n", a, b)
		fmt.Println("Recovered from panic:", r)
	} else {
		fmt.Println("No panic occurred.")
	}
}
