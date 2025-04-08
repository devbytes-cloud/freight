package commit

const (
	// PreCommit is invoked first before the commit process starts. It’s used to inspect or modify the changes being committed. If it exits non-zero, the commit is aborted.
	PreCommit = "pre-commit"
	// PrepareCommitMsg is called after the default commit message is created but before the user is given the chance to edit it. It’s useful for altering the default message.
	PrepareCommitMsg = "prepare-commit-msg"
	// CommitMsg is called after the user has edited the commit message. It’s used to validate or enforce specific commit message formats. If it exits non-zero, the commit is aborted.
	CommitMsg = "commit-msg"
	// PostCommit is invoked after a commit is made. It cannot affect the commit process but can be used for notifications or logging.
	PostCommit = "post-commit"
)

type HookStep struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type Operations struct {
	// PreCommit This hook is invoked first before the commit process starts. It’s used to inspect or modify the changes being committed. If it exits non-zero, the commit is aborted.
	PreCommit []HookStep `json:"pre-commit"`
	// PreCommitMsg This hook is called after the default commit message is created but before the user is given the chance to edit it. It’s useful for altering the default message.
	PrepareCommitMsg []HookStep `json:"prepare-commit-msg"`
	// CommitMsg This hook is called after the user has edited the commit message. It’s used to validate or enforce specific commit message formats. If it exits non-zero, the commit is aborted.
	CommitMsg []HookStep `json:"commit-msg"`
	// CommitMsgPass This hook is
	CommitMsgPass []HookStep `json:"commit-msg-pass"`
	// PostCommit This hook is invoked after a commit is made. It cannot affect the commit process but can be used for notifications or logging.
	PostCommit []HookStep `json:"post-commit"`
}
