package main

import (
	f "fmt"

	cli_tool "github.com/zquan315/learning_golang/w4-cli-file-testing/day1-cli"
	file "github.com/zquan315/learning_golang/w4-cli-file-testing/day2-file"
)

func main() {
	var choice int

	for {
		f.Println("⭐ Welcome to the CLI - File - Testing Learning Program! ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰  1. Day 1: CLI Tool")
		f.Println("☰  2. Day 2: File processing")

		f.Print("Please choose a day to run: ")
		f.Scan(&choice)
		switch choice {
		case 0:
			f.Println("👋 Goodbye! See you next time!")
			return
		case 1:
			f.Println("------Day 1: CLI Tool------")
			cli_tool.RunCLI()
		case 2:
			f.Println("------Day 2: File Processing------")
			file.RunFile()
		default:
			f.Println("❌ Invalid choice, please try again.\n------------")
		}

	}

}
