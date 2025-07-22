package main

import (
	f "fmt"

	"github.com/zquan315/learning_golang/w2-pointer-interface/day5_6_miniProject"
)

func main() {
	var choice int
	for {
		f.Println("⭐ ----- Menu Management ----- ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰ 1. Manage Students")
		f.Println("☰ 2. Manage Banking")
		f.Println("☰ 3. Run basic Interface")
		f.Println("☰ 4. Run advanced Interface")
		f.Println("☰ 5. Run Mini Project")
		f.Print("Enter your choice: ")
		f.Scan(&choice)

		switch choice {
		case 1:
			f.Println("----- Day 1: Student Management -----")
			RunStudentManagement()
		case 2:
			f.Println("----- Day 2: Banking -----")
			RunBanking()
		case 3:
			f.Println("----- Day 3-4: Basic Interface Example -----")
			RunInterface()
		case 4:
			f.Println("----- Day 3-4: Advanced Interface Example -----")
			RunAdvancedInterface()
		case 5:
			f.Println("----- Day 5-6: Mini Project -----")
			day5_6_miniProject.RunMiniProject()
		case 0:
			f.Println("Exiting the program.")
			return
		default:
			f.Println("Invalid choice! Please try again.")
		}
	}
}
