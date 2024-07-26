package blueprint

import (
	"fmt"

	"github.com/devbytes-cloud/freight/internal/blueprint/templates"
	"github.com/devbytes-cloud/freight/internal/validate"
)

var (
	base          = ".git/hooks"
	preCommitPath = fmt.Sprintf("%s/pre-commit", base)
	commitMsgPath = fmt.Sprintf("%s/commit-msg", base)
)

const (
	PreCommit = "PreCommit"
	CommitMsg = "CommitMsg"
)

// NewPreCommit creates a blueprint for the pre-commit script
func NewPreCommit() (*BluePrint, error) {
	dir, err := validate.CurrentWD()
	if err != nil {
		return nil, err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: dir,
		Type: PreCommit,
	}

	return NewBluePrint("pre-commit", preCommitPath, templates.PreHookTmpl, path), nil
}

// NewCommitMsg ...
func NewCommitMsg() (*BluePrint, error) {
	dir, err := validate.CurrentWD()
	if err != nil {
		return nil, err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: dir,
		Type: CommitMsg,
	}

	return NewBluePrint("commit-message", commitMsgPath, templates.PreHookTmpl, path), nil
}
