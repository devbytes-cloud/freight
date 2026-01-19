package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pterm/pterm"

	"github.com/devbytes-cloud/freight/internal/config"
	"github.com/devbytes-cloud/freight/internal/githooks"
)

func main() {
	pterm.EnableDebugMessages()
	pterm.Debug.MessageStyle = pterm.NewStyle(pterm.FgMagenta)
	hookType := os.Args[1]

	printCurrentHook(hookType)

	file, err := os.Open("railcar.json")
	if err != nil {
		pterm.Fatal.Println(err)
	}
	defer func() {
		_ = file.Close()
	}()

	byt, err := io.ReadAll(file)
	if err != nil {
		pterm.Fatal.Println(err)
	}

	var cfg config.Config
	if err := json.Unmarshal(byt, &cfg); err != nil {
		pterm.Fatal.Println(err)
	}

	switch hookType {
	case githooks.CommitMsg:
		commitMsg, err := os.ReadFile(os.Args[2])
		if err != nil {
			pterm.Fatal.Println("Error reading commit message file: ", err)
		}
		if len(cfg.RailCar.CommitOperations.CommitMsg) != 0 {
			run(cfg.RailCar.CommitOperations.CommitMsg, string(commitMsg))
		}

	case githooks.PreCommit:
		if len(cfg.RailCar.CommitOperations.PreCommit) != 0 {
			run(cfg.RailCar.CommitOperations.PreCommit, "")
		}

	case githooks.PrepareCommitMsg:
		commitMsg, err := os.ReadFile(os.Args[2])
		if err != nil {
			pterm.Fatal.Println("Error reading commit message file: ", err)
		}
		if len(cfg.RailCar.CommitOperations.PrepareCommitMsg) != 0 {
			run(cfg.RailCar.CommitOperations.PrepareCommitMsg, string(commitMsg))
		}
	case githooks.PostCommit:
		if len(cfg.RailCar.CommitOperations.PostCommit) != 0 {
			run(cfg.RailCar.CommitOperations.PostCommit, "")
		}
	case githooks.PostCheckout:
		if len(cfg.RailCar.CheckoutOperations.PostCheckout) != 0 {
			run(cfg.RailCar.CheckoutOperations.PostCheckout, "")
		}
	default:
		pterm.Warning.Printfln("Unknown hook type: %s", hookType)
	}

	// Read the commit message from the file

	//•	Commit Hooks Order: pre-commit → prepare-commit-msg → commit-msg → post-commit
	//•	Merge Hooks Order: pre-merge-commit → post-merge
	//•	Rebase Hooks Order: pre-rebase → post-rewrite
	//•	Push Hooks Order: pre-push → update (server-side) → post-update (server-side) → post-receive (server-side)
	//•	Checkout Hooks Order: pre-checkout → post-checkout
	//•	Applypatch Hooks Order: applypatch-msg → pre-applypatch → post-applypatch
}

// run executes a sequence of hook steps with the provided hook data.
func run(data []config.HookStep, hookData string) {
	for _, v := range data {
		pterm.Info.Println(fmt.Sprintf("Name: %s", v.Name))

		cmd := exec.Command("sh", "-c", v.Command)
		if hookData != "" {

			pterm.Debug.Println(fmt.Sprintf("Hook Data: %s", hookData))

			if strings.Contains(v.Command, "{HOOK_INPUT}") {
				hookData = strings.TrimSuffix(hookData, "\n")
				newStr := strings.ReplaceAll(v.Command, "{HOOK_INPUT}", hookData)
				cmd = exec.Command("sh", "-c", newStr)
			} else {
				cmd = exec.Command("sh", "-c", fmt.Sprintf("%s %s", v.Command, hookData))
			}
		}
		pterm.Debug.Println(fmt.Sprintf("Command: %s", cmd.String()))

		cmd.Stdout, cmd.Stderr = &ptermWriter{printFunc: ptermInfoPrintWrapper}, &ptermWriter{printFunc: ptermInfoErrorWrapper}
		if err := cmd.Run(); err != nil {
			pterm.Error.Printfln("Error executing command: %s", err.Error())
			os.Exit(1)
		}
	}
}

// printCurrentHook prints the name of the Git hook currently being executed.
func printCurrentHook(hook string) {
	pterm.DefaultSection.Println("Running hook:", hook)
}

// ptermWriter is an io.Writer that routes its output through pterm.
type ptermWriter struct {
	printFunc func(...any)
}

func ptermInfoPrintWrapper(a ...any) {
	pterm.Info.Println(a...)
}

func ptermInfoErrorWrapper(a ...any) {
	pterm.Error.Println(a...)
}

func (w *ptermWriter) Write(p []byte) (int, error) {
	w.printFunc(string(p))
	return len(p), nil
}
