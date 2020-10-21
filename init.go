package main

import (
	"os"
)

var debugMode bool

func init() {
	if os.Getenv("SLV_DEBUG") != "" {
		debugMode = true
	}
}
