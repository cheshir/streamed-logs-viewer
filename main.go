package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/cheshir/streamed-logs-viewer/internal/app"
)

func main() {
	logger, err := newLogger()
	if err != nil {
		fmt.Println("Failed to start: ", err.Error())

		return
	}

	app := &cli.App{
		Name:  "Streamed logs viewer",
		Usage: "It's like a piped grep but you don't need to restart process to change search criteria.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "buffer-size",
				Value: 100,
				Usage: "how many messages will be stored for processing",
			},
		},
		Action: func(c *cli.Context) error {
			a, err := app.New(app.Config{
				NumberOfBufferedMessages: c.Int("buffer-size"),
			})

			if err != nil {
				return errors.Wrap(err, "failed to init app")
			}

			return a.Run()
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		logger.Error("App crashed", err)
		fmt.Println("Application crashed:", err.Error())

		return
	}
}
