package cmd

import (
	"fmt"
	"strconv"

	"wanto/cli_app/model"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a todo",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		model.Database.Delete(&model.Todos{}, id)
		fmt.Println("Deleted todo with ID:", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
