package termenv

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// Test prompt when user input is true
func TestPromptInputYes(t *testing.T) {
	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := Prompt(strings.NewReader("y\n"), "Yes or No?")

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

			confirmed := Prompt(strings.NewReader(input), "Yes or No?")
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

	confirmed := Prompt(strings.NewReader("maybe\ny\n"), "Yes or No?")
	if !confirmed {
		t.Fatal("expected y after invalid input to confirm")
	}
}

func silenceLogger(t *testing.T, output *bytes.Buffer) {
	t.Helper()

	StandardLogger.SetOutput(output)
	t.Cleanup(func() {
		StandardLogger.SetOutput(os.Stderr)
	})
}
