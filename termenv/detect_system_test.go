//go:build linux && amd64

package termenv

import (
	"os"
	"testing"
)

func TestGetUserHomePath(t *testing.T) {
	homePath, err := GetUserHomePath()
	if err != nil {
		t.Fatalf("GetUserHomePath() returned unexpected error: %v", err)
	}

	if homePath == "" {
		t.Error("GetUserHomePath() returned an empty string")
	}

	expectedEnv := os.Getenv("HOME")
	if expectedEnv != "" && homePath != expectedEnv {
		t.Errorf("Expected home directory %q, but got %q", expectedEnv, homePath)
	}
}
