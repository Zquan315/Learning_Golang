package file

import (
	json "encoding/json"
	f "fmt"
	o "os"
	t "time"
)

func text() {
	f.Println("Reading file README.md....")
	t.Sleep(2 * t.Second) // Simulate a delay for reading the file
	content, err := o.ReadFile(`D:\TO_CONG_QUAN\Learning_Golang\README.md`)
	if err != nil {
		f.Println("Error reading file:", err)
		return
	}
	f.Println("File README.md's content:\n", string(content))
	f.Println("Waiting for 2 seconds before writing a file...")
	t.Sleep(2 * t.Second)
	f.Print("Please enter the content you want to write to the file: ")
	var input string
	f.Scan(&input)
	o.WriteFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\text.txt`, []byte(input), 0644)
	f.Println("File text.txt has been written successfully with the content:", input)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func json_sample() {
	p := Person{Name: "Quan", Age: 21}
	f.Println("Converting struct to JSON...")
	jsondata, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		f.Println("Error converting to JSON:", err)
		return
	}
	f.Println("Wrieting JSON data to file person.json...")
	o.WriteFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\person.json`, jsondata, 0644)
	f.Println("Read the JSON data from file person.json...")
	json_content, err := o.ReadFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\person.json`)
	if err != nil {
		f.Println("Error reading JSON file:", err)
		return
	}
	f.Println("JSON data read from file:\n", string(json_content))
	var temp_Person Person
	err = json.Unmarshal(jsondata, &temp_Person)
	if err != nil {
		f.Println("Error reading JSON data:", err)
		return
	}
	f.Println("Temp Person:", temp_Person)
}

type task struct {
	Name   string `json:"Name"`
	Status bool   `json:"Status"`
}

func task_sample() {
	task := make([]task, 3)
	for i := 0; i < len(task); i++ {
		f.Print("Please enter the name and status (0/1) of task ", i+1, ": ")
		var status string
		f.Scan(&task[i].Name, &status)
		if status == "1" {
			task[i].Status = true
		} else {
			task[i].Status = false
		}
	}
	f.Println("Converting task to JSON...")
	file, err1 := o.OpenFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\task.json`, o.O_APPEND|o.O_CREATE|o.O_WRONLY, 0664)
	if err1 != nil {
		f.Println("Error open file")
		return
	}
	defer file.Close()
	for _, value := range task {
		jsondata, err := json.MarshalIndent(value, "", "  ")
		if err != nil {
			f.Println("Error converting to JSON:", err)
			return
		}
		_, err = file.Write(append(jsondata, '\n', '\n'))
		if err != nil {
			f.Println("Error writing JSON to file:", err)
			return
		}
	}

	json_content, err := o.ReadFile(`D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\day2-file\task.json`)
	if err != nil {
		f.Println("Error reading JSON file:", err)
		return
	}
	f.Println("JSON data read from file:\n", string(json_content))
}
func RunFile() {
	//text()
	//json_sample()
	task_sample()
	f.Println("⭐ File processing completed successfully!")
}
