/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	f "fmt"
	o "os"
	json "encoding/json"
	t "time"
	"github.com/spf13/cobra"
	i "github.com/zquan315/learning_golang/w4-cli-file-testing/todo/init"
)

var taskName string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creta a task",
	Long:  "This command is init to create a task",
	Run: func(cmd *cobra.Command, args []string) {
		if taskName == "" {
			f.Println("The flag name is required value ")
			return;
		}
		fileJsonPath := `D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\todo\Tasks.json`
		dataTask := []i.Task{}
		jsoncontent, err1 := o.ReadFile(fileJsonPath)
		if err1 == nil {
			err1 = json.Unmarshal(jsoncontent, &dataTask)
			if err1 != nil {
				f.Println("❌ Error in decode!")
				return
			}
		}
		newTask := i.Task{
			ID: t.Now().Format("20060102150405"),
			Name: taskName,
			Status: false,
		}
		dataTask = append(dataTask, newTask)

		jsondata, err := json.MarshalIndent(dataTask,"", "\t")
		if err != nil {
			f.Println("Error in convert Task to json data")
			return
		}
		
		err = o.WriteFile(fileJsonPath, jsondata, 0644)
		if err != nil {
			f.Println("Error in WriteFile in json")
			return
		}

		f.Printf("✅ Create a task succesfully with ID: %s and Name: %s", newTask.ID, taskName)
	},
}

func init() {
	createCmd.Flags().StringVarP(&taskName, "name", "n","", "Name of task is required")
	createCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(createCmd)
}
