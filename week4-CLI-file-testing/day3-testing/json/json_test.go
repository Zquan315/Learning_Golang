package json

import (
	f "fmt"
	"testing"
)

func TestJson(t *testing.T) {
	tasks, err := DecodeJsontoTask()
	if err != nil {
		f.Println("Error in Decode!")
	}

	if len(tasks) != 3 {
		t.Errorf("Error! Len of tasks is %d", len(tasks))
	}

	exp := []Task{
		{"Play", true},
		{"Sleep", false},
		{"Learn", true},
	}

	for i, v := range tasks {
		if v != exp[i] {
			t.Errorf("Failed! %v != %v", v, exp[i])
		}
	}
}
