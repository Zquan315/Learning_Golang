package json

import (
	json "encoding/json"
	f "fmt"
	o "os"
)

type Task struct {
	Name   string `json:"Name"`
	Status bool   `json:"Status"`
}

func DecodeJsontoTask() ([]Task, error) {
	jsondata, err := o.ReadFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\task.json`)
	if err != nil {
		f.Println("Error in Reading file json!")
		return nil, err
	}

	var tasks []Task
	err2 := json.Unmarshal(jsondata, &tasks)
	if err2 != nil {
		f.Println("Error in Encoding json data!")
		return nil, err
	}
	return tasks, err
}
