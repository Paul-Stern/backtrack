package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backtrack",
	Short: "backtrack is a simple post-factum tracker",
	Run: func(cmd *cobra.Command, args []string) {
		// log.Info().Interface("args", args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Caller().Err(err)
	}
}
