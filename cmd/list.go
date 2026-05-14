package cmd

import (
	"github.com/spf13/cobra"
)

var ListTasks = &cobra.Command{
	Use:   "ls",
	Short: "list all task",
	Long:  "list all the task you have added",
	RunE: func(cmd *cobra.Command, args []string) error {
		completion, _ := cmd.Flags().GetBool("completed")
		pending, _ := cmd.Flags().GetBool("pending")
		all, _ := cmd.Flags().GetBool("all")
		priority, _ := cmd.Flags().GetString("priority")
		tag, _ := cmd.Flags().GetString("tag")
		sortBy, _ := cmd.Flags().GetString("sort")
		limit, _ := cmd.Flags().GetInt64("limit")
		orderType, _ := cmd.Flags().GetString("OrderType")

		err := NewTodo().ListTodo(completion, pending, all, priority, tag, sortBy, limit, orderType)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	ListTasks.Flags().BoolP("completed", "c", false, "List only completed tasks")
	ListTasks.Flags().BoolP("pending", "p", true, "List only pending tasks")
	ListTasks.Flags().BoolP("all", "a", false, "List all tasks")
	ListTasks.Flags().StringP("priority", "P", "midium", "List tasks with a specific priority")
	ListTasks.Flags().StringP("tag", "t", "", "List tasks with a specific tag")
	ListTasks.Flags().StringP("sort", "s", "id", "Sort the tasks by the given field")
	ListTasks.Flags().Int64P("limit", "l", 10, "Limit the number of tasks returned")
	ListTasks.Flags().StringP("OrderType", "o", "DESC", "write ASC for ascending order and DESC for descending order")
	rootCmd.AddCommand(ListTasks)
}
