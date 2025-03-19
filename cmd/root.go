package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "CLI Todo app",
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
