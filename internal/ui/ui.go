package ui

import (
	"github.com/rivo/tview"
)

type App struct {
	view  *tview.Application
	views *Views
}

func New() *App {
	app := tview.NewApplication()

	return &App{
		view:  app,
		views: newViews(app),
	}
}

func (a *App) Run() error {
	a.initHandlers()

	return a.view.Run()
}

func (a *App) Output() *Output {
	return &Output{
		view: a.views.Logs,
	}
}
