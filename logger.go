package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const logFile = "./app.log"

func newLogger() (*logrus.Logger, error) {
	logger := logrus.New()

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, errors.Wrapf(err, `failed to open log file "%s"`, logFile)
	}

	logger.SetOutput(f)

	if debugMode {
		logger.Level = logrus.DebugLevel
	}

	return logger, nil
}
