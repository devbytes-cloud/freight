package validate

import (
	"fmt"
	"os"
	"path/filepath"
)

var (
	// gitDir is the relative path to the Git directory within a repository.
	gitDir = ".git"
	// gitHookDir is the relative path to the Git hooks directory inside .git.
	// It is constructed using filepath.Join to ensure OS-independent path separators.
	gitHookDir = filepath.Join(gitDir, "hooks")
)

// GitDirs checks to see if there is a valid .git file in the repo
func GitDirs() error {
	if _, err := os.Stat(gitDir); err != nil {
		return fmt.Errorf(".git directory missing: %w", err)
	}

	if _, err := os.Stat(gitHookDir); err != nil {
		return fmt.Errorf(".git/hooks directory missing: %w", err)
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
