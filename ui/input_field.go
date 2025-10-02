package ui

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TaskInputField struct {
	*tview.InputField

	app *App
}

func NewTaskInputField(app *App) *TaskInputField {
	f := TaskInputField{
		InputField: tview.NewInputField(),
	}

	f.SetLabel("Enter task message: ").
		SetBorder(true).
		SetTitle("Input")
	f.SetDoneFunc(func(key tcell.Key) {
		t := model.NewTask(f.InputField.GetText())
		t.Add()
		app.UpdateTasks()
		app.Main.SwitchToPage("list")
	})
	return &f
}
