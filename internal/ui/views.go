package ui

import (
	"github.com/rivo/tview"
)

type View string

const (
	statusBarWidth = 18

	helpPage = "help"
	mainPage = "main"
)

func newViews(viewApp *tview.Application) *Views {
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

	helpView := tview.NewTextView().
		SetWordWrap(true)

	_, _ = statusBar.Write([]byte("   .* json wrap   ")) // todo move it

	pages := tview.NewPages()
	pages.AddPage(mainPage, container, false, true)
	pages.AddPage(helpPage, helpView, false, false)

	viewApp.SetRoot(pages, true)

	return &Views{
		Container:    container,
		Help:         helpView,
		Logs:         logsView,
		Pages:        pages,
		Search:       searchField,
		StatusBar:    statusBar,
		TopContainer: topContainer,
	}
}

type Views struct {
	Container    *tview.Flex
	Help         *tview.TextView
	Logs         *tview.TextView
	Pages        *tview.Pages
	Search       *tview.InputField
	StatusBar    *tview.TextView
	TopContainer *tview.Flex
}

func (v *Views) ShowSearch() {
	v.TopContainer.
		RemoveItem(v.StatusBar).
		AddItem(v.Search, 0, 1, true).
		AddItem(v.StatusBar, statusBarWidth, 0, false)
}

func (v *Views) HideSearch() {
	v.TopContainer.
		RemoveItem(v.Search).
		ResizeItem(v.StatusBar, 0, 1)
}

func (v *Views) ShowHelp() {
	v.Pages.HidePage(mainPage).ShowPage(helpPage)
}

func (v *Views) HideHelp() {
	v.Pages.HidePage(helpPage).ShowPage(mainPage)
}
