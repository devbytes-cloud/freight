package githooks

import (
	"path/filepath"
)

// hooksBaseDir is the base directory for Git hooks
const hooksBaseDir = ".git/hooks"

// GitHooks contains an array of GitHook
type GitHooks struct {
	Commit   []GitHook
	Checkout []GitHook
}

// GitHook represents a Git hook with its name, path, and template
type GitHook struct {
	// Name of the Git hook (also the type)
	Name string
	// Path to the particular Git hook
	Path string
	// Template of the file contents for the Git hook
	Template string
}

// NewGitHooks returns a pointer to a GitHooks instance with commit hooks initialized
func NewGitHooks() *GitHooks {
	return &GitHooks{
		Commit:   generateHooks(getCommitHook()),
		Checkout: generateHooks(getCheckoutHooks()),
	}
}

// generateHooks creates a slice of GitHook instances based on the provided hook names.
// Each GitHook includes its name, path, and a template for the hook script.
func generateHooks(hook []string) []GitHook {
	hooks := make([]GitHook, len(hook))
	for i, hookName := range hook {
		hooks[i] = GitHook{
			Name:     hookName,
			Path:     hookPath(hookName),
			Template: gitHookTemplate,
		}
	}
	return hooks
}

// hookPath configures the Git hook path for a given hook type
func hookPath(hookType string) string {
	return filepath.Join(hooksBaseDir, hookType)
}
