package githooks

import (
	"fmt"

	"github.com/devbytes-cloud/freight/internal/blueprint/templates"
	"github.com/devbytes-cloud/freight/internal/githooks/commit"
)

// hooksBaseDir is the base directory for Git hooks
const hooksBaseDir = ".git/hooks"

// GitHooks contains an array of GitHook
type GitHooks struct {
	Commit []GitHook
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
		Commit: generateCommitHooks(),
	}
}

// hookPath configures the Git hook path for a given hook type
func hookPath(hookType string) string {
	return fmt.Sprintf("%s/%s", hooksBaseDir, hookType)
}

func generateCommitHooks() []GitHook {
	commitHookNames := []string{
		commit.PreCommit,
		commit.PrepareCommitMsg,
		commit.CommitMsg,
		commit.PostCommit,
	}

	hooks := make([]GitHook, len(commitHookNames))
	for i, hookName := range commitHookNames {
		hooks[i] = GitHook{
			Name:     hookName,
			Path:     hookPath(hookName),
			Template: templates.PreHookTmpl, // Consider using a different template if needed
		}
	}
	return hooks
}
