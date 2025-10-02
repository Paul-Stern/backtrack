package ui

import "github.com/rivo/tview"

type TaskTable struct {
	*tview.Table
	app *App
}

func NewTaskTable(app *App) *TaskTable {
	return &TaskTable{
		Table: tview.NewTable(),
		app:   app,
	}
}

func (t *TaskTable) Update() {

}
