/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	f "fmt"
	json "encoding/json"
	o "os"
	i "github.com/zquan315/learning_golang/w4-cli-file-testing/todo/init"
	"github.com/spf13/cobra"
)
var idFlag string
// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark task to done",
	Long:  "Use done to mark task to done",
	Run: func(cmd *cobra.Command, args []string) {
		fileJsonPath := "tasks.json"

		if idFlag == "" {
			f.Println("Value of ID flag is not empty")
			return
		}

		jsoncontent, err := o.ReadFile(fileJsonPath)
		if err != nil {
			f.Println(err)
			return
		}

		listTask := make([]i.Task,0)
		err = json.Unmarshal(jsoncontent, &listTask)
		if err != nil {
			f.Println(err)
			return
		}

		check := false
		for i, v := range listTask {
			if v.ID == idFlag {
				listTask[i].Status = true
				check = true
				break
			}
		}

		if !check {
			f.Printf("❌ %s is not exists", idFlag)
			return 
		}

		jsondata, error := json.MarshalIndent(listTask, "", "\t")
		if error != nil {
			f.Println(error)
			return
		}

		error = o.WriteFile(fileJsonPath, jsondata, 0644)
		if error != nil {
			f.Println(error)
			return 
		}

		f.Printf("✅ Mark a task succesfully with ID: %s", idFlag)
	},	
}

func init() {
	doneCmd.Flags().StringVarP(&idFlag, "id", "","", "ID is required")
	doneCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(doneCmd)
}
