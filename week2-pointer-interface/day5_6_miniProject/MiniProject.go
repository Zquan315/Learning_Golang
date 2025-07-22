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
		f.Println("☰ 4. Delete a Student")
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
		case 3:
			f.Println("----- Find a Student by ID -----")
			var id string
			f.Print("Enter Student ID: ")
			f.Scan(&id)
			st := findStudentByID(Students, id)
			f.Println("Result\n", st)
		case 4:
			f.Println("----- Delete Student -----")
			var id string
			f.Print("Enter Student ID to delete: ")
			f.Scan(&id)
			Students = delStudent(Students, id)
		case 5:
			f.Println("----- Sort Students by GPA -----")
			sortByGPA(Students)
		default:
			f.Println("Invalid choice! Please try again.")
		}
	}
}
