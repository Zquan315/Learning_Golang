package main

import f "fmt"

func main() {
	var choice int
	for {
		f.Println("⭐ ----- Menu Management ----- ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰ 1. Manage Students")
		f.Println("☰ 2. Manage Banking")
		f.Print("Enter your choice: ")
		f.Scan(&choice)

		switch choice {
		case 1:
			f.Println("----- Day 1: Student Management -----")
			RunStudentManagement()
		case 2:
			f.Println("----- Day 2: Banking -----")
			RunBanking()
		case 0:
			f.Println("Exiting the program.")
			return
		default:
			f.Println("Invalid choice! Please try again.")
		}
	}
}
