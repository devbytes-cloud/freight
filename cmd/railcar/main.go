package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/devbytes-cloud/freight/internal/blueprint"
)

type RailCar struct {
	CommitOperations CommitOperations `json:"commit-operations"`
}

type CommitOperations struct {
	// PreCommit This hook is invoked first before the commit process starts. It’s used to inspect or modify the changes being committed. If it exits non-zero, the commit is aborted.
	PreCommit map[string]string `json:"pre-commit"`
	// PreCommitMsg This hook is called after the default commit message is created but before the user is given the chance to edit it. It’s useful for altering the default message.
	PrepareCommitMsg map[string]string `json:"prepare-commit-msg"`
	// CommitMsg This hook is called after the user has edited the commit message. It’s used to validate or enforce specific commit message formats. If it exits non-zero, the commit is aborted.
	CommitMsg map[string]string `json:"commit-msg"`
	// CommitMsgPass This hook is
	CommitMsgPass map[string]string `json:"commit-msg-pass"`
	// PostCommit This hook is invoked after a commit is made. It cannot affect the commit process but can be used for notifications or logging.
	PostCommit map[string]string `json:"post-commit"`
}
type Config struct {
	RailCar RailCar `json:"carriage"`
}

func main() {
	hookType := os.Args[1]

	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	byt, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(byt, &config); err != nil {
		panic(err)
	}

	switch hookType {
	case blueprint.CommitMsg:
		//commitMsg, err := ioutil.ReadFile(os.Args[2])
		//if err != nil {
		//	fmt.Println("Error reading commit message file:", err)
		//	os.Exit(1)
		//}

		// Process the commit message
		fmt.Println("Commit message is:", os.Args[2])
		if len(config.RailCar.CommitOperations.CommitMsg) != 0 {
			run(config.RailCar.CommitOperations.CommitMsg, os.Args[2])
		}

	case blueprint.PreCommit:
		if len(config.RailCar.CommitOperations.PreCommit) != 0 {
			run(config.RailCar.CommitOperations.PreCommit, "")
		}

	case blueprint.PrepareCommitMsg:
		if len(config.RailCar.CommitOperations.PrepareCommitMsg) != 0 {
			run(config.RailCar.CommitOperations.PrepareCommitMsg, "")
		}
	case blueprint.PostCommit:
		if len(config.RailCar.CommitOperations.PostCommit) != 0 {
			run(config.RailCar.CommitOperations.PostCommit, "")
		}
	}

	// Read the commit message from the file

	//•	Commit Hooks Order: pre-commit → prepare-commit-msg → commit-msg → post-commit
	//•	Merge Hooks Order: pre-merge-commit → post-merge
	//•	Rebase Hooks Order: pre-rebase → post-rewrite
	//•	Push Hooks Order: pre-push → update (server-side) → post-update (server-side) → post-receive (server-side)
	//•	Checkout Hooks Order: pre-checkout → post-checkout
	//•	Applypatch Hooks Order: applypatch-msg → pre-applypatch → post-applypatch
}

func run(data map[string]string, hookData string) {
	for k, v := range data {

		fmt.Println(fmt.Sprintf("RUNNING :: %s", k))

		cmd := exec.Command("sh", "-c", v)
		if hookData != "" {
			cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", v, hookData))
		}
		fmt.Println(cmd.String())
		cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
			os.Exit(1)
		}
	}
}
