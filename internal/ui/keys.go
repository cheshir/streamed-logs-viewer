package ui

import (
	"github.com/gdamore/tcell/v2"
)

const (
	KeyEsq              = Key(tcell.KeyEsc)
	KeyEnter            = Key(tcell.KeyEnter)
	KeyCtrlC            = Key(tcell.KeyCtrlC)
	KeyQuestionMark Key = '?'
	KeySlash        Key = '/'
	KeyB            Key = 'b'
	KeyR            Key = 'r'
	KeyM            Key = 'm'
	KeyI            Key = 'i'
)

type Key byte

func (k Key) String() string {
	switch k {
	case KeyEsq:
		return "Escape"
	case KeyEnter:
		return "Enter"
	case KeyCtrlC:
		return "Ctrl-c"
	default:
		return string(k)
	}
}
