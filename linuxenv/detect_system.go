package linuxenv

import (
	"fmt"
	"os"
	"runtime"

	"github.com/nsymx/scoopbucket/internal/logger"
)

func DetectOS() {
	if runtime.GOOS != "linux" || runtime.GOARCH != "amd64" {
		logger.Standard.Error("Unsupported operating system detected")
		logger.Standard.Fatal("Support available for linux-amd64 only")
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
