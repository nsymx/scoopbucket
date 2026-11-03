package styles

import (
	"charm.land/lipgloss/v2"
)

// Define colors
var (
	colorText  = lipgloss.Color("#FFFFFF")
	colorAlpha = lipgloss.Color("#1016cf")
)

// Basic text style
var BasicTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorText))

// Alpha text style
var AlphaTextStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorAlpha))

// Greeter style
var GreeterStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorText)).
	BorderStyle(lipgloss.HiddenBorder()).
	BorderForeground((colorAlpha)).
	MarginTop(1).
	MarginLeft(1)

// Simple enum style
var EnumeratorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground((colorAlpha))
