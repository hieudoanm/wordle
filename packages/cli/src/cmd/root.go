package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wordle",
	Short: "wordle CLI application (games tools)",
	Long:  `The wordle CLI application is a comprehensive backend utility belonging to the games suite of tools.

Use this root executable to manage configuring, running, and interacting with all wordle-related operations securely and efficiently from your terminal.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}
