package cmd

import (
	"fmt"
	"os"
	"tsk/color"

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
		errorString := color.ApplyStyle("Oops. An error occurred while executing tsk '%s'\n", color.Red, color.Bold)
		fmt.Fprintf(os.Stderr, errorString, err)
		os.Exit(1)
	}
}
