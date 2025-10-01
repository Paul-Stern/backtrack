package commands

import (
	"fmt"

	"github.com/Paul-Stern/backtrack/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists journal entries",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(model.GetTasks())
	},
}
