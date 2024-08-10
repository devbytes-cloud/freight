package githooks

import (
	"fmt"

	"github.com/devbytes-cloud/freight/internal/blueprint/templates"
	"github.com/devbytes-cloud/freight/internal/githooks/commit"
)

// GitHooks contains an array of GitHook
type GitHooks struct {
	Commit []GitHook
}

// GitHook represents a githook with its name, path, and template
type GitHook struct {
	// Name of the githook (also the type)
	Name string
	// Path to the particular githook
	Path string
	// Template of the file contents for the githook
	Template string
}

// NewGitHooks returns a pointer to a GitHooks
func NewGitHooks() *GitHooks {
	return &GitHooks{
		Commit: []GitHook{
			{
				Name:     commit.PreCommit,
				Path:     hookPath(commit.PreCommit),
				Template: templates.PreHookTmpl,
			},
			{
				Name:     commit.PrepareCommitMsg,
				Path:     hookPath(commit.PrepareCommitMsg),
				Template: templates.PreHookTmpl,
			},
			{
				Name:     commit.CommitMsg,
				Path:     hookPath(commit.CommitMsg),
				Template: templates.PreHookTmpl,
			},
			{
				Name:     commit.PostCommit,
				Path:     hookPath(commit.PostCommit),
				Template: templates.PreHookTmpl,
			},
		},
	}
}

// hookPath configures the githook path with the particular type
func hookPath(hookType string) string {
	base := ".git/hooks"
	return fmt.Sprintf("%s/%s", base, hookType)
}
