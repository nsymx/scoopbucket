package termenv

import (
	"os"

	"github.com/charmbracelet/log"
)

var StandardLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: true,
	Prefix:          ":",
})

var CompactLogger = log.NewWithOptions(os.Stderr, log.Options{
	ReportTimestamp: false,
	Prefix:          ":",
})
