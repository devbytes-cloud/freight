package validate

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/devbytes-cloud/freight/internal/githooks"
	"github.com/pterm/pterm"
)

var (
	// gitDir is the relative path to the Git directory within a repository.
	gitDir = ".git"
	// gitHookDir is the relative path to the Git hooks directory inside .git.
	// It is constructed using filepath.Join to ensure OS-independent path separators.
	gitHookDir = filepath.Join(gitDir, "hooks")
)

// GitDirs verifies if the current directory contains a valid .git directory and the hooks subdirectory.
func GitDirs() error {
	if _, err := os.Stat(gitDir); err != nil {
		return fmt.Errorf(".git directory missing: %w", err)
	}

	if _, err := os.Stat(gitHookDir); err != nil {
		return fmt.Errorf(".git/hooks directory missing: %w", err)
	}
	return nil
}

// CurrentWD returns the absolute path of the current working directory.
func CurrentWD() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return dir, nil
}

// GitHooks validates the provided allow hooks and returns a map of valid hooks.
func GitHooks(allow []string) (map[string]struct{}, error) {
	allowedGitHooks := githooks.AllowedGitHooks()
	if len(allow) == 0 {
		pterm.Debug.Println("No hooks provided, using default allowed hooks")
		return allowedGitHooks, nil
	}

	inputHooks := map[string]struct{}{}
	var invalidHooks []string
	for _, v := range allow {
		if _, ok := allowedGitHooks[v]; !ok {
			invalidHooks = append(invalidHooks, v)
			continue
		}
		inputHooks[v] = struct{}{}
	}

	if len(invalidHooks) > 0 {
		return nil, fmt.Errorf("invalid hook types: %s", strings.Join(invalidHooks, ", "))
	}

	return inputHooks, nil
}
