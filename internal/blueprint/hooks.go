package blueprint

import (
	"github.com/devbytes-cloud/freight/internal/githooks"

	"github.com/devbytes-cloud/freight/internal/validate"
)

func NewGitHook(gh *githooks.GitHook) (*BluePrint, error) {
	dir, err := validate.CurrentWD()
	if err != nil {
		return nil, err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: dir,
		Type: gh.Name,
	}
	return NewBluePrint(gh.Name, gh.Path, gh.Template, path), nil
}
