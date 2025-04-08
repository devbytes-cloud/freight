package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/devbytes-cloud/freight/internal/githooks/commit"
)

type RailCar struct {
	CommitOperations commit.Operations `json:"commit-operations"`
}

type Config struct {
	RailCar RailCar `json:"railcar"`
}

func main() {
	hookType := os.Args[1]

	file, err := os.Open("railcar.json")
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

	fmt.Println("hook type is ", hookType)

	switch hookType {
	case commit.CommitMsg:
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

	case commit.PreCommit:
		fmt.Println(config.RailCar)
		if len(config.RailCar.CommitOperations.PreCommit) != 0 {
			fmt.Println("whores")
			run(config.RailCar.CommitOperations.PreCommit, "")
		}

	case commit.PrepareCommitMsg:
		if len(config.RailCar.CommitOperations.PrepareCommitMsg) != 0 {
			run(config.RailCar.CommitOperations.PrepareCommitMsg, "")
		}
	case commit.PostCommit:
		if len(config.RailCar.CommitOperations.PostCommit) != 0 {
			run(config.RailCar.CommitOperations.PostCommit, "")
		}
	default:
		fmt.Println("couldnt find hook type which was")
		fmt.Println(hookType)
	}

	// Read the commit message from the file

	//•	Commit Hooks Order: pre-commit → prepare-commit-msg → commit-msg → post-commit
	//•	Merge Hooks Order: pre-merge-commit → post-merge
	//•	Rebase Hooks Order: pre-rebase → post-rewrite
	//•	Push Hooks Order: pre-push → update (server-side) → post-update (server-side) → post-receive (server-side)
	//•	Checkout Hooks Order: pre-checkout → post-checkout
	//•	Applypatch Hooks Order: applypatch-msg → pre-applypatch → post-applypatch
}

func run(data []commit.HookStep, hookData string) {
	for _, v := range data {

		fmt.Println(fmt.Sprintf("RUNNING :: %s", v.Name))
		cmd := exec.Command("sh", "-c", v.Command)
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
