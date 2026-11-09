package termenv

import (
	"fmt"
	"os"
	"runtime"
)

func DetectOS() {
	if runtime.GOOS != "linux" || runtime.GOARCH != "amd64" {
		StandardLogger.Error("Unsupported operating system detected")
		StandardLogger.Fatal("Support available for linux-amd64 only")
	}
}

func GetUserHomePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	info, err := os.Stat(homePath)
	if err != nil {
		return "", err
	}
	if !info.IsDir() {
		return "", fmt.Errorf("not a directory: %s", homePath)
	}

	return homePath, nil
}
