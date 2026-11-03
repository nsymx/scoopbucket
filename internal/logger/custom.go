package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var BasicLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})
