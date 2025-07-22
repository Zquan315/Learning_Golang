package day5_6_miniProject

import f "fmt"

func sortByGPA(students []*Student) {
	st := make([]*Student, len(students))
	copy(st, students)
	for i := 0; i < len(st)-1; i++ {
		for j := i + 1; j < len(st); j++ {
			if st[i].GPA < st[j].GPA {
				t := st[i]
				st[i] = st[j]
				st[j] = t
			}
		}
	}
	f.Println("Students sorted by GPA in descending order:")
	showAllStudents(st)
}
