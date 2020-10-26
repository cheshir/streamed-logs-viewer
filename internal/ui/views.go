package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type View string

const (
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

	helpView := newHelpView()
	_, _ = statusBar.Write([]byte("   .*/mi json wrap   ")) // todo move it

	pages := tview.NewPages()
	pages.AddPage(mainPage, container, false, true)
	pages.AddPage(helpPage, newHelpContainer(helpView), false, false)

	viewApp.SetRoot(pages, true)

	return &Views{
		app:          viewApp,
		Container:    container,
		Help:         helpView,
		Logs:         logsView,
		Pages:        pages,
		Search:       searchField,
		StatusBar:    statusBar,
		TopContainer: topContainer,
	}
}

type helpRecord struct {
	Action, Description string
}

var helpData = []helpRecord{
	{"ESC", "Return to command mode"},
	{"?", "Help"},
	{"/", "Search"},
	{"f", "Search"},
	{"i", "Enable ignore case regexp mode"},
	{"j", "Beautify JSON"},
	{"m", "Enable multiline regexp mode"},
	{"r", "Enable regexp (golang syntax)"},
	{"w", "Wrap lines"},
	{"Ctrl-c", "Exit"},
}

const padding = "    "

func newHelpView() *tview.Table {
	t := tview.NewTable()
	row := 0

	for _, record := range helpData {
		actionCell := tview.NewTableCell(record.Action + padding).
			SetTextColor(tcell.ColorDodgerBlue).
			SetAttributes(tcell.AttrBold)

		descriptionCell := tview.NewTableCell(record.Description)

		t.SetCell(row, 0, actionCell)
		t.SetCell(row, 1, descriptionCell)
		row++
	}

	return t
}

func newHelpContainer(helpView tview.Primitive) *tview.Flex {
	helpVerticalContainer := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(helpView, 0, 1, true).
		AddItem(tview.NewBox(), 0, 1, false)

	helpContainer := tview.NewFlex().
		SetFullScreen(true).
		AddItem(tview.NewBox(), 0, 1, false).
		AddItem(helpVerticalContainer, 0, 1, false).
		AddItem(tview.NewBox(), 0, 1, false)

	return helpContainer
}

type Views struct {
	app          *tview.Application
	Container    *tview.Flex
	Help         *tview.Table
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
		AddItem(v.StatusBar, len(v.StatusBar.GetText(true)), 0, false)

	v.app.SetFocus(v.Search)
}

func (v *Views) HideSearch() {
	v.TopContainer.
		RemoveItem(v.Search).
		ResizeItem(v.StatusBar, 0, 1)

	v.Logs.ScrollToEnd()
	v.app.SetFocus(v.Logs)
}

func (v *Views) ShowHelp() {
	v.Pages.HidePage(mainPage).ShowPage(helpPage)
	v.app.SetFocus(v.Help)
}

func (v *Views) HideHelp() {
	v.Pages.HidePage(helpPage).ShowPage(mainPage)
	v.Logs.ScrollToEnd()
	v.app.SetFocus(v.Logs)
}
