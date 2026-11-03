package main

import (
	"flag"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/nsymx/scoopbucket/internal/styles"
	"github.com/nsymx/scoopbucket/linuxenv"
)

var (
	name       = "scoopbucket"
	version    = "dev"
	enumerator = "󰟓 "
)

func greeter() {
	lipgloss.Println(
		styles.GreeterStyle.Render(
			styles.EnumeratorStyle.Render(enumerator),
			styles.BasicTextStyle.Render(
				"Welcome to"), styles.AlphaTextStyle.Render(name), "!"))
}

func main() {
	linuxenv.DetectOS()
	versionFlag := flag.Bool("version", false, "Show latest version")

	flag.Parse()
	if *versionFlag {
		lipgloss.Println(styles.BasicTextStyle.Render(
			name), "-", strings.TrimSpace(version))
		os.Exit(0)
	}

	linuxenv.ClearScreen()
	greeter()
}
