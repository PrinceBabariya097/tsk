package cmd

import (
	"github.com/spf13/cobra"
)

var ListTasks = &cobra.Command{
	Use:   "ls",
	Short: "list all task",
	Long:  "list all the task you have added",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := NewTodo().ListTodo()
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ListTasks)
}
