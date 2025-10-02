package commands

import (
	"fmt"

	"github.com/Paul-Stern/backtrack/model"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "lists journal entries",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		// if len(args) > 0 {
		// 	cmd.ValidateArgs(args)
		// }
		fmt.Println(model.GetTasks(10).String())
	},
}

func init() {
	listCmd.AddCommand(todayCmd)
	listCmd.AddCommand(sinceCmd)
	listCmd.AddCommand(allCmd)
}

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "lists journal entries for today",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(model.Today().String())
	},
}

var sinceCmd = &cobra.Command{
	Use:     "since",
	Short:   "lists journal entries since date",
	Example: "backtrack since 2022-01-14\nbacktrack since $(date -d '20220114' -I)",
	Args:    cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		t := model.TryParseTime(args[0])
		fmt.Println(model.Since(t).String())
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "lists journal entries for today",
	Args:  cobra.MatchAll(cobra.ExactArgs(0)),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(model.GetAll().String())
	},
}
