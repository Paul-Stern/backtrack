package view

import "github.com/Paul-Stern/backtrack/ui"

type App struct {
	*ui.App
	Content *PageStack
}

func NewApp() *App {
	return &App{
		App:     ui.NewApp(),
		Content: NewPageStack(),
	}
}

func (a *App) QueueUpdateDraw(f func()) {
	if a.Application == nil {
		return
	}
	go func() {
		a.Application.QueueUpdateDraw(f)
	}()
}

func (a *App) Run() error {
	func() {
		a.Main.SwitchToPage("list")
	}()

	return a.Application.Run()
}

func (a *App) UpdateTasks() {
	a.Main.Pages.GetPage("list").(*ui.TaskList).Update()
}
