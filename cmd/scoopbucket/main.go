package main

import (
	"flag"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/nsymx/scoopbucket/termenv"
)

var (
	name       = "scoopbucket"
	version    = "dev"
	enumerator = "󰟓 "
)

func greeter() {
	lipgloss.Println(
		termenv.GreeterStyle.Render(
			termenv.EnumeratorStyle.Render(enumerator),
			termenv.BasicTextStyle.Render(
				"Welcome to"), termenv.AlphaTextStyle.Render(name), "!"))
}

func main() {
	termenv.DetectOS()
	versionFlag := flag.Bool("version", false, "Show latest version")

	flag.Parse()
	if *versionFlag {
		lipgloss.Println(termenv.BasicTextStyle.Render(
			name), "-", strings.TrimSpace(version))
		os.Exit(0)
	}

	termenv.ClearScreen()
	greeter()
}
