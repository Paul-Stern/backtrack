package commands

import (
	"github.com/Paul-Stern/backtrack/ui"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

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
	// app := tview.NewApplication()
	app := ui.NewApp()
	l := ui.NewTaskView(app)
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := app.SetRoot(l, true).SetFocus(l).Run(); err != nil {
		log.Panic().Err(err)
	}
}
