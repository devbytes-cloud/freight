package validate

import (
	"fmt"
	"os"
)

var (
	basePath = ".git"
	hookPath = fmt.Sprintf("%s/hooks", basePath)
)

// GitDirExists checks to see if there is a valid .git file in the repo
func GitDirExists() error {
	if _, err := os.Stat(basePath); err != nil {
		return err
	}

	if _, err := os.Stat(hookPath); err != nil {
		return err
	}
	return nil
}

// CurrentWD returns the current directory that you are in
func CurrentWD() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}
