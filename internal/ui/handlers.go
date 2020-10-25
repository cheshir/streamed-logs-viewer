package ui

import (
	"github.com/gdamore/tcell/v2"
)

func (a *App) initHandlers() {
	logs := a.Logs()
	search := a.Search()

	logs.SetChangedFunc(func() {
		a.view.Draw()
	})

	logs.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEnter:
			a.showSearch()
			a.view.SetFocus(search)
		}

		switch event.Rune() {
		case '?':
			_, _ = logs.Write([]byte("Caught ????\n"))
		}

		return event
	})

	search.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			a.hideSearch()
			logs.ScrollToEnd()
			a.view.SetFocus(logs)
		}

		_, _ = logs.Write([]byte(">>> " + search.GetText() + "\n"))
	})
}
