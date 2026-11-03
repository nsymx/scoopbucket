package linuxenv

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/nsymx/scoopbucket/internal/logger"
	"github.com/nsymx/scoopbucket/internal/styles"
)

func ClearScreen() {
	_, err := fmt.Print("\033[2J\033[H")
	if err == nil {
		return
	}

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Run()
}

func ShellPrompt(input io.Reader, NewPrompt string) bool {
	reader := bufio.NewReader(input)

	for {
		logger.CompactLogger.Print(styles.BasicTextStyle.Render(NewPrompt))

		userInput, _ := reader.ReadString('\n')
		userInput = strings.ToLower(strings.TrimSpace(userInput))

		switch userInput {
		case "y":
			return true
		case "n", "":
			return false
		default:
			fmt.Println()
			logger.CompactLogger.Warn("invalid input, please type y or n...")
			fmt.Println()
		}
	}
}
