package ui

import (
	"github.com/gdamore/tcell/v2"
)

func (a *App) initHandlers() {
	a.views.Logs.SetChangedFunc(func() {
		a.view.Draw()
	})

	a.views.Logs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			a.views.ShowSearch()
		}

		switch event.Rune() {
		case '?':
			a.views.ShowHelp()
		}

		return event
	})

	a.views.Help.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			a.views.HideHelp()
		}

		return event
	})

	a.views.Search.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			a.views.HideSearch()
		}

		_, _ = a.views.Logs.Write([]byte(">>> " + a.views.Search.GetText() + "\n")) // todo
	})
}
