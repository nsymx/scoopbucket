package termenv

import (
	"bytes"
	"errors"
	"os"
	"strings"
	"testing"
)

func TestPromptInputYes(t *testing.T) {
	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := Prompt(strings.NewReader("y\n"), "Yes or No?")

	if !confirmed {
		t.Fatal("expected y to confirm")
	}
}

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

func TestPromptInvalidInput(t *testing.T) {
	var output bytes.Buffer

	silenceLogger(t, &output)

	confirmed := Prompt(strings.NewReader("maybe\ny\n"), "Yes or No?")
	if !confirmed {
		t.Fatal("expected y after invalid input to confirm")
	}
}

func TestPromptInputAtEOF(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{name: "yes", input: "y", expected: true},
		{name: "invalid", input: "maybe", expected: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var output bytes.Buffer
			silenceLogger(t, &output)

			confirmed := Prompt(strings.NewReader(test.input), "Yes or No?")
			if confirmed != test.expected {
				t.Fatalf("Prompt(%q) = %t, want %t", test.input, confirmed, test.expected)
			}
		})
	}
}

func TestPromptReadError(t *testing.T) {
	var output bytes.Buffer
	silenceLogger(t, &output)

	confirmed := Prompt(errorReader{}, "Yes or No?")
	if confirmed {
		t.Fatal("expected a read error to cancel")
	}
}

type errorReader struct{}

func (errorReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}

func silenceLogger(t *testing.T, output *bytes.Buffer) {
	t.Helper()

	StandardLogger.SetOutput(output)
	t.Cleanup(func() {
		StandardLogger.SetOutput(os.Stderr)
	})
}
