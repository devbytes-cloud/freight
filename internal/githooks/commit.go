package githooks

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

// getCommitHook returns a slice of strings representing the names of commit-related Git hooks.
// These hooks include pre-commit, prepare-commit-msg, commit-msg, and post-commit.
func getCommitHook() []string {
	return []string{
		PreCommit,
		PrepareCommitMsg,
		CommitMsg,
		PostCommit,
	}
}
