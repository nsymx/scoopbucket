package termenv

import (
	"charm.land/lipgloss/v2"
)

// Basic text style
var BasicTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((ColorText))

// Alpha text style
var AlphaTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((ColorAlpha))

// Greeter style
var GreeterStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((ColorText)).
	BorderStyle(lipgloss.HiddenBorder()).
	BorderForeground((ColorAlpha)).
	MarginTop(1).
	MarginLeft(1)

// Simple enum style
var EnumeratorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((ColorAlpha))
