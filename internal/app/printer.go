package app

import (
	"io"

	"github.com/cheshir/streamed-logs-viewer/internal/ui"
)

type Printer struct {
	in         io.Reader
	out        *ui.Output
	formatters []formatter
}
