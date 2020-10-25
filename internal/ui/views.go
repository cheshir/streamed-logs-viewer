package ui

import (
	"github.com/rivo/tview"
)

type View string

const (
	ContainerView    View = "container"
	LogsView         View = "text"
	SearchView       View = "search"
	StatusBarView    View = "status_bar"
	TopContainerView View = "top_container"
)

const (
	statusBarWidth = 18
)

type Views map[View]tview.Primitive

func initViews(viewApp *tview.Application) Views {
	logsView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(true)

	logsView.SetBorder(true)

	searchField := tview.NewInputField().
		SetPlaceholder("Search")

	statusBar := tview.NewTextView().
		SetTextAlign(tview.AlignRight).
		SetDynamicColors(true)

	topContainer := tview.NewFlex().
		AddItem(statusBar, 0, 1, false)

	container := tview.NewFlex().
		SetFullScreen(true).
		SetDirection(tview.FlexRow).
		AddItem(topContainer, 1, 1, false).
		AddItem(logsView, 0, 1, true)

	_, _ = statusBar.Write([]byte("   .* json wrap   ")) // todo move it

	pages := tview.NewPages()
	pages.AddPage(mainPage, container, false, true)

	viewApp.SetRoot(pages, true)

	return Views{
		ContainerView:    container,
		LogsView:         logsView,
		SearchView:       searchField,
		StatusBarView:    statusBar,
		TopContainerView: topContainer,
	}
}

func (a *App) showSearch() {
	a.TopContainer().
		RemoveItem(a.views[StatusBarView]).
		AddItem(a.views[SearchView], 0, 1, true).
		AddItem(a.views[StatusBarView], statusBarWidth, 0, false)
}

func (a App) hideSearch() {
	a.TopContainer().
		RemoveItem(a.views[SearchView]).
		ResizeItem(a.views[StatusBarView], 0, 1)
}
