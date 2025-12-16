package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListTasks = &cobra.Command{
	Use:   "ls",
	Short: "list all task",
	Long:  "list all the task you have added",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Printing...")
		err := NewTodo().ListTodo()
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(ListTasks)
}
