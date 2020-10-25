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
			a.view.SetFocus(a.views.Search)
		}

		switch event.Rune() {
		case '?':
			a.views.ShowHelp()
			a.view.SetFocus(a.views.Help)
		}

		return event
	})

	a.views.Help.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEsc:
			a.views.HideHelp()
			a.views.Logs.ScrollToEnd()
			a.view.SetFocus(a.views.Logs)
		}

		return event
	})

	a.views.Search.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			a.views.HideSearch()
			a.views.Logs.ScrollToEnd()
			a.view.SetFocus(a.views.Logs)
		}

		_, _ = a.views.Logs.Write([]byte(">>> " + a.views.Search.GetText() + "\n")) // todo
	})
}
