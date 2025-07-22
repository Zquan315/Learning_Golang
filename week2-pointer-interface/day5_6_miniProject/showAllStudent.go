package day5_6_miniProject

import (
	f "fmt"
	"strconv"
)

func showAllStudents(students []*Student) {
	if len(students) == 0 {
		f.Println("No students available.")
		return
	}
	for i, student := range students {
		gpa, _ := strconv.ParseFloat(student.GPA, 64)
		f.Printf("%d. %s\t%s\t%s\t%.2f\n", i+1, student.ID, student.Name, student.Gender, gpa)
	}
}
