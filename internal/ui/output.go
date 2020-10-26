package ui

import (
	"github.com/rivo/tview"
)

type Output struct {
	view *tview.TextView
}

func (o Output) Write(data []byte) (int, error) {
	return o.view.Write(data)
}

func (o Output) Clear() {
	o.view.Clear()
}
