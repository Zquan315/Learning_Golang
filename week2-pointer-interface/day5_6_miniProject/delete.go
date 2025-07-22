package day5_6_miniProject

import f "fmt"

func delStudent(students []*Student, id string) []*Student {
	for i, st := range students {
		if st.ID == id {
			students = append(students[:i], students[i+1:]...)
			f.Println("Deleted Student with ID:", id, "successfully.")
			return students
		}
	}
	f.Println("Student not found with ID:", id)
	return students
}
