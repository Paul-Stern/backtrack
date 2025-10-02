package commands

import "github.com/spf13/cobra"

var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "runs the mainloop",
	Run: func(cmd *cobra.Command, args []string) {
		MainLoop()
	},
}

func init() {
	rootCmd.AddCommand(loopCmd)
}

func MainLoop() {
	// main loop
}
