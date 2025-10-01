package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "backtrack",
	Short: "backtrack is a simple post-factum tracker",
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Caller().Err(err)
	}
}
