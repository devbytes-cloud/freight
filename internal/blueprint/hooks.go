package blueprint

import (
	"fmt"

	"github.com/devbytes-cloud/freight/internal/blueprint/templates"
	"github.com/devbytes-cloud/freight/internal/validate"
)

var (
	base                 = ".git/hooks"
	preCommitPath        = fmt.Sprintf("%s/pre-commit", base)
	prepareCommitMsgPath = fmt.Sprintf("%s/prepare-commit-msg", base)
	commitMsgPath        = fmt.Sprintf("%s/commit-msg", base)
	postCommitPath       = fmt.Sprintf("%s/post-commit", base)
)

const (
	PreCommit        = "PreCommit"
	PrepareCommitMsg = "PrepareCommitMsg"
	CommitMsg        = "CommitMsg"
	PostCommit       = "PostCommit"
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

// NewPrepareCommitMsg ...
func NewPrepareCommitMsg() (*BluePrint, error) {
	dir, err := validate.CurrentWD()
	if err != nil {
		return nil, err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: dir,
		Type: PrepareCommitMsg,
	}

	return NewBluePrint("pre-commit-message", prepareCommitMsgPath, templates.PreHookTmpl, path), nil
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

// NewPostCommit ...
func NewPostCommit() (*BluePrint, error) {
	dir, err := validate.CurrentWD()
	if err != nil {
		return nil, err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: dir,
		Type: PostCommit,
	}

	return NewBluePrint("post-commit", postCommitPath, templates.PreHookTmpl, path), nil
}
