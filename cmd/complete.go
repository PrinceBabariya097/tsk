package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeTask = &cobra.Command{
	Use:   "done",
	Short: "complete task",
	Long:  "complete task that you have done using id",
	RunE: func(cmd *cobra.Command, args []string) error {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("Please enter valid number")
		}
		error := NewTodo().CompleteTodo(int64(num))
		if error != nil {
			errorMessage := errors.New("error while finding index in db")
			return errors.Join(errorMessage, error)
		}
		fmt.Printf("Task is completed for ID: %d", int64(num))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeTask)
}
