package ui

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/rivo/tview"
)

type TaskView struct {
	*tview.List

	Items model.Tasks
}

func NewTaskView(app *App) *TaskView {
	tv := TaskView{
		Items: model.GetTasks(10),
	}
	tv.List = tv.Items.List()
	tv.AddItem("Add", "Add a new task", 'a', func() {
		app.SetRoot(NewTaskInputField(app), true)
	})
	tv.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	return &tv
}
