package view

import "github.com/Paul-Stern/backtrack/ui"

type PageStack struct {
	*ui.Pages

	app *App
}

func NewPageStack() *PageStack {
	return &PageStack{
		Pages: ui.NewPages(),
	}
}
