package cmd

import (
	"fmt"
	"strconv"

	"wanto/cli_app/model"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [id]",
	Short: "Mark a todo as completed",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		var todo model.Todos
		model.Database.First(&todo, id)
		if todo.ID == 0 {
			fmt.Println("Todo not found")
			return
		}

		todo.Done = true
		model.Database.Save(&todo)
		fmt.Println("Marked as completed:", todo.Text)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
