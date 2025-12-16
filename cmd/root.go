package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tsk",
	Short: "tsk is cli tool to store your task",
	Long:  "tsk is cli tool to store your task. you can set reminder of that task",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error occurred while executing tsk '%s'\n", err)
		os.Exit(1)
	}
}
