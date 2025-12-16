package cmd

import "github.com/spf13/cobra"

var addTskCmd = &cobra.Command{
	Use:     "Add",
	Aliases: []string{"add"},
	Short:   "Add task title",
	Long:    "Add a title of your task that you want to add",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "" {
			cmd.PrintErr("please add your task to add it")
		}
		NewTodo().AddTodo(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addTskCmd)
}
