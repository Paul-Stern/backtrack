package ui

import (
	"github.com/Paul-Stern/backtrack/model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rs/zerolog/log"
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
		if err := t.Add(); err != nil {
			log.Error().Err(err).Msg("Failed to add task")
		}
	})
	return &f
}
