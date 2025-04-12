package config

// Config represents the top-level configuration structure.
// It encapsulates the RailCar configuration, which includes commit and checkout operations.
type Config struct {
	// RailCar holds the configuration for Git hook operations.
	RailCar RailCar `json:"config"`
}

// RailCar represents the configuration for Git hook operations.
// It includes both commit-related operations and checkout-related operations.
type RailCar struct {
	// CommitOperations defines the set of hook steps executed during various stages of the commit process.
	CommitOperations CommitOperations `json:"commit-operations"`

	// CheckoutOperations defines the steps executed during checkout-related Git operations.
	CheckoutOperations CheckoutOperation `json:"checkout-operations"`
}

// CommitOperations encapsulates the set of hook steps executed at various stages of the commit process.
type CommitOperations struct {
	// PreCommit This hook is invoked first before the commit process starts. It’s used to inspect or modify the changes being committed. If it exits non-zero, the commit is aborted.
	PreCommit []HookStep `json:"pre-commit"`
	// PreCommitMsg This hook is called after the default commit message is created but before the user is given the chance to edit it. It’s useful for altering the default message.
	PrepareCommitMsg []HookStep `json:"prepare-commit-msg,omitempty"`
	// CommitMsg This hook is called after the user has edited the commit message. It’s used to validate or enforce specific commit message formats. If it exits non-zero, the commit is aborted.
	CommitMsg []HookStep `json:"commit-msg,omitempty"`
	// CommitMsgPass This hook is
	CommitMsgPass []HookStep `json:"commit-msg-pass,omitempty"`
	// PostCommit This hook is invoked after a commit is made. It cannot affect the commit process but can be used for notifications or logging.
	PostCommit []HookStep `json:"post-commit,omitempty"`
}

type CheckoutOperation struct {
	// PostCheckout represents the steps to be executed after a checkout operation is completed.
	PostCheckout []HookStep `json:"post-checkout"`
}

// HookStep represents a single step in the execution sequence of a Git hook.
//
// Each HookStep consists of a name that identifies the step (for example,
// "lint" or "format") and a command that specifies the operation or script
// that should be executed during that step. This struct is typically used to
// build a sequence of actions within a Git hook's workflow.
type HookStep struct {
	// Name holds the identifier for the hook step.
	// It distinguishes this step within the sequence (e.g., "lint", "test", or "format")
	// and is used when serializing/deserializing to JSON.
	Name string `json:"name"`

	// Command specifies the command or script that will be executed for this hook step.
	// It defines the particular operation that should be performed during this step
	// of the Git hook process.
	Command string `json:"command"`
}
