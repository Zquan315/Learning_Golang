/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "mark task to done",
	Long:  "Use done to mark task to done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("done called")
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
