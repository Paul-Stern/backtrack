package ui

import (
	"github.com/rivo/tview"
)

type App struct {
	*tview.Application

	Main  *Pages
	views map[string]tview.Primitive
}

func NewApp() *App {
	a := App{
		Application: tview.NewApplication(),
		Main:        NewPages(),
	}

	a.views = map[string]tview.Primitive{
		"list":  NewTaskList(&a),
		"input": NewTaskInputField(&a),
	}
	a.Main.AddPage("list", a.views["list"], true, true)
	a.Main.AddPage("input", a.views["input"], true, false)
	return &a
}

func (a *App) UpdateTasks() {
	a.Main.GetPage("list").(*TaskList).Update()
}
