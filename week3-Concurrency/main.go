package main

import (
	f "fmt"

	goroutine_channel "github.com/zquan315/learning_golang/w3-concurrency/day1-2-goroutine-channel"
)

func main() {
	var choice int

	for {
		f.Println("⭐ Welcome to the Concurrency Learning Program! ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰  1. Day 1: Goroutine")
		f.Println("☰  2. Day 2: Channel")

		f.Print("Please choose a day to run: ")
		f.Scan(&choice)
		switch choice {
		case 0:
			f.Println("⭐ Exiting the program.")
			return
		case 1:
			f.Println("--- Day 1: Goroutine ---")
			goroutine_channel.RunGoRoutine()
		case 2:
			f.Println("--- Day 2: Basic Channel ---")
			goroutine_channel.RunChannel()
		default:
			f.Println("❌ Invalid choice, please try again.\n------------")
		}

	}

}
