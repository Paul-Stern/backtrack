package view

import "github.com/Paul-Stern/backtrack/ui"

type App struct {
	*ui.App
	Content *PageStack
}

func (a *App) Run() error {
	func() {
		a.Main.SwitchToPage("list")
	}()

	return a.Application.Run()
}
