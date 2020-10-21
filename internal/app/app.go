package app

import (
	"bufio"
	"io"

	"github.com/derailed/tview"
	"github.com/gdamore/tcell"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/cheshir/streamed-logs-viewer/internal/input"
)

type Config struct {
	NumberOfBufferedMessages int
}

type App struct {
	buffer *messageBuffer
	input  *bufio.Reader
}

func New(config Config) (*App, error) {
	in, err := input.New()
	if err != nil {
		return nil, errors.Wrap(err, "input error")
	}

	app := &App{
		buffer: newMessageBuffer(config.NumberOfBufferedMessages),
		input:  in,
	}

	return app, nil
}

func (a *App) Run() error {
	ui := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			ui.Draw()
		})

	textView.SetDoneFunc(func(key tcell.Key) {
		switch key {
		case tcell.KeyEnter:
			_, _ = textView.Write([]byte("<Enter> pressed\n"))
		case tcell.KeyTAB:
			_, _ = textView.Write([]byte("<Tab> pressed\n"))
		}
	})

	textView.SetBorder(true)

	errs := errgroup.Group{}
	errs.Go(func() error {
		for {
			line, err := a.input.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				return errors.Wrap(err, "read input error")
			}

			if _, err := textView.Write(line); err != nil {
				if err == io.EOF {
					break
				}

				return errors.Wrap(err, "write text error")
			}
		}

		return nil
	})

	errs.Go(func() error {
		return ui.
			SetRoot(textView, true).
			SetFocus(textView).
			Run()
	})

	return errs.Wait()
}
