package ui

import (
	"github.com/rivo/tview"
)

type App struct {
	view  *tview.Application
	views Views
}

func New() *App {
	app := tview.NewApplication()

	return &App{
		view:  app,
		views: initViews(app),
	}
}

func (a *App) Run() error {
	a.initHandlers()

	return a.view.Run()
}

func (a *App) Container() *tview.Flex {
	return a.views[ContainerView].(*tview.Flex)
}

func (a *App) Logs() *tview.TextView {
	return a.views[LogsView].(*tview.TextView)
}

func (a *App) Search() *tview.InputField {
	return a.views[SearchView].(*tview.InputField)
}

func (a *App) TopContainer() *tview.Flex {
	return a.views[TopContainerView].(*tview.Flex)
}
