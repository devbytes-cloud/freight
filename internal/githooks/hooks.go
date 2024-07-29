package githooks

import (
	"fmt"

	"github.com/devbytes-cloud/freight/internal/blueprint/templates"

	"github.com/devbytes-cloud/freight/internal/githooks/commit"
)

type GitHooks struct {
	Commit []GitHook
}

type GitHook struct {
	Name     string
	Path     string
	Template string
}

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

func hookPath(hookType string) string {
	base := ".git/hooks"
	return fmt.Sprintf("%s/%s", base, hookType)
}
