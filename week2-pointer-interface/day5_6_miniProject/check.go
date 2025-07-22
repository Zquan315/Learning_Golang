package day5_6_miniProject

import (
	"regexp"
	"strconv"
)

func CheckExists(check map[string]bool, id string) bool {
	if _, ok := check[id]; ok {
		return true
	} else {
		return false
	}
}

func CheckValidName(name string) bool {
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(name)
}

func CheckValidID(id string) bool {
	return regexp.MustCompile(`^[0-9]{8}$`).MatchString(id)
}
func CheckValidGender(gender string) bool {
	if gender == "Male" || gender == "Female" {
		return true
	}
	return false
}
func CheckValidInput(name, id, gender, gpa string) bool {
	if name == "" || id == "" || gender == "" {
		return false
	}
	if valid, err := strconv.ParseFloat(gpa, 64); err != nil || valid < 0.0 || valid > 10.0 {
		return false
	}
	if !CheckValidName(name) || !CheckValidID(id) || !CheckValidGender(gender) {
		return false
	}
	return true
}
