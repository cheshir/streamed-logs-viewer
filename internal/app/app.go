package app

import (
	"bufio"
	"io"

	"github.com/pkg/errors"

	"github.com/cheshir/streamed-logs-viewer/internal/input"
	"github.com/cheshir/streamed-logs-viewer/internal/ui"
)

type Config struct {
	NumberOfBufferedMessages int
}

type App struct {
	buffer *messageBuffer
	input  *bufio.Reader
	ui     *ui.App
}

func New(config Config) (*App, error) {
	in, err := input.New()
	if err != nil {
		return nil, errors.Wrap(err, "input error")
	}

	app := &App{
		buffer: newMessageBuffer(config.NumberOfBufferedMessages),
		input:  in,
		ui:     ui.New(),
	}

	return app, nil
}

func (a *App) Run() error {
	errs := make(chan error, 1)

	go func() {
		errs <- func() error {
			logsBlock := a.ui.Views().Logs

			for {
				line, err := a.input.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						break
					}

					return errors.Wrap(err, "read input error")
				}

				if _, err := logsBlock.Write(line); err != nil {
					if err == io.EOF {
						break
					}

					return errors.Wrap(err, "write text error")
				}
			}

			return nil
		}()
	}()

	go func() {
		errs <- a.ui.Run()
	}()

	return <-errs
}
