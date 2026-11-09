package termenv

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Prompt(input io.Reader, NewPrompt string) bool {
	reader := bufio.NewReader(input)

	for {
		CompactLogger.Print(BasicTextStyle.Render(NewPrompt))

		userInput, _ := reader.ReadString('\n')
		userInput = strings.ToLower(strings.TrimSpace(userInput))

		switch userInput {
		case "y":
			return true
		case "n", "":
			return false
		default:
			fmt.Println()
			CompactLogger.Warn("invalid input, please type y or n...")
			fmt.Println()
		}
	}
}
