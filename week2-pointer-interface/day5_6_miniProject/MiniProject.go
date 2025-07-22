package day5_6_miniProject

import (
	f "fmt"
)

func RunMiniProject() {
	Students := make([]*Student, 0)
	check := make(map[string]bool)
	var choice int
	for {
		f.Println("⭐ ----- Mini Project: Student Management ----- ⭐")
		f.Println("❌ 0. Exit")
		f.Println("☰ 1. Add Student")
		f.Println("☰ 2. List all Students")
		f.Println("☰ 3. Find a Student by ID")
		f.Println("☰ 3. Delete Student")
		f.Println("☰ 4. Update Student")
		f.Println("☰ 5. Sort Students by GPA")
		f.Print("Enter your choice: ")
		f.Scan(&choice)

		switch choice {
		case 0:
			f.Println("Exiting the Mini Project.")
			return
		case 1:
			f.Println("----- Add Student -----")
			Students = addStudent(Students, check)
		case 2:
			f.Println("----- List of Students -----")
			showAllStudents(Students)
		default:
			f.Println("Invalid choice! Please try again.")
		}
	}
}
