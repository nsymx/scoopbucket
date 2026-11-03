package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

var BasicLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

var CompactLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: false,
	Prefix:          ":",
})
