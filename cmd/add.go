package cmd

import (
	"fmt"

	"wanto/cli_app/model"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "new todos",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := model.Todos{Text: args[0]}
		model.Database.Create(&todo)
		fmt.Println("added: ", todo.Text)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
