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
var statusFlag string
var allFlag bool
var checkStatusFlag bool
var stFlag bool
// listCmd represents the list command
var listCmd = &cobra.Command {
	Use:   "list",
	Short: "list all task",
	Long:  "Use list command to list all task",
	Run: func(cmd *cobra.Command, args []string) {
		fileJsonPath := `D:\TO_CONG_QUAN\Learning_Golang\week4-CLI-file-testing\todo\Tasks.json`
		if !allFlag {
			f.Println("Flag -all or -a is required")
			return
		}

		listTask := make([]i.Task, 0)
		jsoncontent, err := o.ReadFile(fileJsonPath)
		if err != nil {
			f.Println("Error in Reading a file json!")
			return
		}

		err = json.Unmarshal(jsoncontent, &listTask)
		if err != nil {
			f.Println("Error in Unmarshal json to Task!")
			return
		}

		f.Println("┌───────────────┬──────────────────────┬────────────┐")
		f.Printf("│ %-13s │ %-20s │ %-10s │\n", "ID", "Name", "Status")
		f.Println("├───────────────┼──────────────────────┼────────────┤")
		if len(listTask) == 0 {
			f.Println("List is empty")
			return
		}

		if statusFlag != "" {
			if statusFlag != "true" && statusFlag != "false" {
				f.Println("Status is only true or false value")
				return
			} 
			checkStatusFlag = true
			if statusFlag == "true" {
				stFlag = true
			} else {
				stFlag = false
			}
		}

		if checkStatusFlag == true {
			for _, v := range listTask {
				if stFlag == v.Status {
					f.Printf("│ %-8s │ %-19s │ %-10v │\n", v.ID, v.Name, v.Status)
				}
			}
		} else {
			for _, v := range listTask {
				f.Printf("│ %-8s │ %-19s │ %-10v │\n", v.ID, v.Name, v.Status)
			}
		}
		// In đường kẻ ngang dưới cùng
		f.Println("└───────────────┴──────────────────────┴────────────┘")
	},
}
func init() {
	listCmd.Flags().BoolVarP(&allFlag,"all","a", false , "Flag all is required")
	listCmd.MarkFlagRequired("all")

	listCmd.Flags().StringVarP(&statusFlag, "status", "s", "", "List task via status")
	rootCmd.AddCommand(listCmd)

}