package ui

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/rivo/tview"
)

type TaskView struct {
	*tview.List
	app *App

	Items model.Tasks
}

func NewTaskView(app *App) *TaskView {
	tv := TaskView{
		Items: model.GetTasks(10),
		app:   app,
	}
	tv.List = tv.Items.List()
	tv.AddItem("Add", "Add a new task", 'a', func() {
		app.Main.SwitchToPage("input")
	})
	tv.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})
	return &tv
}

func (t *TaskView) Update() {
	t.Clear()
	t.Items = model.GetTasks(10)
	t.List = t.Items.List()
	t.AddMenu()
}

func (t *TaskView) AddMenu() {
	t.AddItem("Add", "Add a new task", 'a', func() {
		t.app.Main.SwitchToPage("input")
	})
	t.AddItem("Quit", "Press to exit", 'q', func() {
		t.app.Stop()
	})
}

func (t *TaskView) QuitItem() {
	t.AddItem("Quit", "Press to exit", 'q', func() {
		t.app.Stop()
	})
}
func (t *TaskView) AddMenuItem() {
	t.AddItem("Add", "Add a new task", 'a', func() {
		t.app.Main.SwitchToPage("input")
	})
}
