package cmd

import (
	"errors"
	"fmt"
	"tsk/color"

	"github.com/spf13/cobra"
)

var addTskCmd = &cobra.Command{
	Use:     "Add",
	Aliases: []string{"add"},
	Short:   "Add task title",
	Long:    "Add a title of your task that you want to add",
	RunE: func(cmd *cobra.Command, args []string) error {
		priority, err := cmd.Flags().GetString("priority")
		tag, err := cmd.Flags().GetString("tag")
		if err != nil {
			cmd.PrintErr(err.Error())
		}
		if len(args) == 0 {
			return errors.New("Task title is required")
		}
		if len(args) > 1 {
			return errors.New("Only one task title is allowed")
		}
		todo, todoError := NewTodo().AddTodo(args[0], priority, tag)

		if todoError != nil {
			return todoError
		}

		printString := color.ApplyStyle("Success: your task '%v' is added successfully", color.Green, color.Bold, color.Italic)
		fmt.Printf(printString, todo.Name)

		return nil
	},
}

func init() {
	addTskCmd.Flags().StringP("priority", "p", "medium", "Set the priority of the task (low, medium, high)")
	addTskCmd.Flags().StringP("tag", "t", "", "Set a tag for the task")
	rootCmd.AddCommand(addTskCmd)
}
