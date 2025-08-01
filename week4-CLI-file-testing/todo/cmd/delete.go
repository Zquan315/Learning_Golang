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
var idFlagDel string
// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long: "Delete a task in list using command delete",
	Run: func(cmd *cobra.Command, args []string) {
		fileJsonPath := "tasks.json"

		if idFlagDel == "" {
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
			if v.ID == idFlagDel {
				listTask = append(listTask[:i], listTask[i+1:]...)
				check = true
				break
			}
		}

		if !check {
			f.Printf("❌ %s is not exists", idFlagDel)
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

		f.Printf("✅ Delete a task succesfully with ID: %s", idFlagDel)

	},
}

func init() {
	deleteCmd.Flags().StringVarP(&idFlagDel, "id", "","", "ID is required")
	deleteCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(deleteCmd)

}
