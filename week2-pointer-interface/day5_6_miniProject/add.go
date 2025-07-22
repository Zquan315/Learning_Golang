package day5_6_miniProject

import f "fmt"

func addStudent(s []*Student, check map[string]bool) []*Student {
	for {
		var name, id, gender, gpa string
		f.Print("Enter Student Name (Only contain a-z or A-Z)): ")
		f.Scan(&name)
		f.Print("Enter Student ID (8 digits ): ")
		f.Scan(&id)
		f.Print("Enter gender (Male or Female): ")
		f.Scan(&gender)
		f.Print("Enter GPA (0.0 - 10.0): ")
		f.Scan(&gpa)
		if !CheckValidInput(name, id, gender, gpa) {
			f.Println("Invalid input! Please try again.")
			continue
		}
		if CheckExists(check, id) {
			f.Println("Student with this ID already exists. Please try again.")
			continue
		}
		st := Student{
			Name:   name,
			ID:     id,
			Gender: gender,
			GPA:    gpa,
		}
		s = append(s, &st)
		f.Println("⭐ Student added successfully! ⭐")
		return s
	}
}
