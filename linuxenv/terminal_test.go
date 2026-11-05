package linuxenv

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/nsymx/scoopbucket/internal/logger"
)

func TestClearScreen(t *testing.T) {
	ClearScreen()
}

// Test prompt when user input is true
func TestPromptInputYes(t *testing.T) {
	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := ShellPrompt(strings.NewReader("y\n"), "Yes or No?")

	if !confirmed {
		t.Fatal("expected y to confirm")
	}
}

// Test prompt when user input is empty or false
func TestPromptEmptyInputNo(t *testing.T) {
	tests := map[string]string{
		"no":    "n\n",
		"empty": "\n",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			var output bytes.Buffer
			silenceLogger(t, &output)

			confirmed := ShellPrompt(strings.NewReader(input), "Yes or No?")
			if confirmed {
				t.Fatalf("expected %q to cancel", input)
			}
		})
	}
}

// Test prompt when user input is invalid
func TestPromptInvalidInput(t *testing.T) {
	var output bytes.Buffer

	silenceLogger(t, &output)

	confirmed := ShellPrompt(strings.NewReader("maybe\ny\n"), "Yes or No?")
	if !confirmed {
		t.Fatal("expected y after invalid input to confirm")
	}
}

func silenceLogger(t *testing.T, output *bytes.Buffer) {
	t.Helper()

	logger.Standard.SetOutput(output)
	t.Cleanup(func() {
		logger.Standard.SetOutput(os.Stderr)
	})
}
