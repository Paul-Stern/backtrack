package commands

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/rivo/tview"
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
	app := tview.NewApplication()
	l := model.GetTasks(10).List()
	l.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	// box := tview.NewBox().SetBorder(true).SetTitle("Hello, world!")
	if err := app.SetRoot(l, true).SetFocus(l).Run(); err != nil {
		log.Panic().Err(err)
	}
}
