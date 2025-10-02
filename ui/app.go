package ui

import "github.com/rivo/tview"

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
		"list":  NewTaskView(&a),
		"input": NewTaskInputField(&a),
	}
	return &a
}
