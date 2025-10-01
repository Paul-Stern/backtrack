package commands

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds journal entry",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	Run: func(cmd *cobra.Command, args []string) {
		// log.Info().Interface("args", args)

		t := model.NewTask(args[0])
		if err := t.Add(); err != nil {
			log.Fatal().Err(err)
		}
	},
}
