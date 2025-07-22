package day5_6_miniProject

import f "fmt"

func findStudentByID(students []*Student, id string) Student {
	for _, st := range students {
		if id == st.ID {
			return *st
		}
	}
	f.Println("Student not found with ID:", id)
	return Student{
		ID:     "",
		Name:   "",
		GPA:    "",
		Gender: "",
	}
}
