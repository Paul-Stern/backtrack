package commands

import (
	"strings"

	"github.com/Paul-Stern/backtrack/model"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds journal entry",
	Run: func(cmd *cobra.Command, args []string) {
		Add(args)
	},
}

func Add(args []string) {
	b := strings.Builder{}
	for i, arg := range args {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(arg)
	}
	msg := b.String()

	t := model.NewTask(msg)
	if err := t.Add(); err != nil {
		log.Fatal().Err(err).Msg("Failed to add task")
	}

}
