package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/nsymx/scoopbucket/termenv"
)

var (
	name    = "scoopbucket"
	version = "dev"
	url     = "github.com/nsymx/scoopbucket"
)

func greeter() {
	fmt.Println(termenv.GreeterStyle.Render(
		"Welcome to", name, "!"))
}

func main() {
	termenv.DetectOS()
	termenv.ClearScreen()

	helpFlag := flag.Bool("h", false, "Show help")
	versionFlag := flag.Bool("v", false, "Print latest version")

	flag.Parse()

	if *helpFlag {
		fmt.Println("", name, "help:")
		fmt.Println()

		fmt.Println("", "-h", "-", "Show this message")
		fmt.Println("", "-v", "-", "Print latest version")
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println(termenv.BasicTextStyle.Render(
			"", name), "-", strings.TrimSpace(version))

		os.Exit(0)
	}

	greeter()
}
