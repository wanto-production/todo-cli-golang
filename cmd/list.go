package cmd

import (
	"fmt"
	"os"

	"wanto/cli_app/model"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Run: func(cmd *cobra.Command, args []string) {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Text", "Status"})
		var todos []model.Todos
		model.Database.Find(&todos)

		fmt.Println("Todo List:")
		for _, todo := range todos {
			status := "❌"
			if todo.Done {
				status = "✅"
			}
			table.Append([]string{fmt.Sprintf("%d", todo.ID), todo.Text, status})
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
