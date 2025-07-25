package main

import (
	f "fmt"

	goroutine_channel "github.com/zquan315/learning_golang/w3-concurrency/day1-2-goroutine-channel"
	channel_immediate "github.com/zquan315/learning_golang/w3-concurrency/day3-4-channel-immediate"
	sync "github.com/zquan315/learning_golang/w3-concurrency/day5-sync"
	mini_project "github.com/zquan315/learning_golang/w3-concurrency/day6-7-miniProject"
)

func main() {
	var choice int

	for {
		f.Println("⭐ Welcome to the Concurrency Learning Program! ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰  1. Day 1: Goroutine")
		f.Println("☰  2. Day 2: Basic Channel")
		f.Println("☰  3. Day 3-4: Immediate Channel")
		f.Println("☰  4. Day 5: Sync for Goroutine")
		f.Println("☰  5. Day 6-7: Mini Project - Calculator a Array of Numbers")

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
			goroutine_channel.RunBasicChannel()
		case 3:
			f.Println("--- Day 3-4: Immediate Channel ---")
			channel_immediate.RunImmediateChannel()
		case 4:
			f.Println("--- Day 5: Sync for Goroutine ---")
			sync.RunSync()
		case 5:
			f.Println("--- Day 6-7: Mini Project - Calculator a Array of Numbers ---")
			mini_project.RunMiniProject()
		default:
			f.Println("❌ Invalid choice, please try again.\n------------")
		}

	}

}
