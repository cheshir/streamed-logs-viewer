package input

import (
	"bufio"
	"errors"
	"os"
)

func New() (*bufio.Reader, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		return nil, errors.New("input must be piped")
	}

	return bufio.NewReader(os.Stdin), nil
}
