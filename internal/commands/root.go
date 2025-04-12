package commands

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/devbytes-cloud/freight/internal/blueprint"
	"github.com/devbytes-cloud/freight/internal/config"
	"github.com/devbytes-cloud/freight/internal/embed"
	"github.com/devbytes-cloud/freight/internal/githooks"
	"github.com/devbytes-cloud/freight/internal/validate"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// Execute runs the root command and handles any errors that occur during execution.
func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NewRootCmd creates and returns the root command for the CLI application.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "init",
		Short: "init",
		Long:  `init`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := validate.GitDirs(); err != nil {
				cmd.PrintErrln(err)
			}
			if err := setupHooks(); err != nil {
				cmd.PrintErrln(err)
			}

			configForce, err := cmd.Flags().GetBool("config-force")
			if err != nil {
				cmd.PrintErrln(err)
			}

			if err := setupConfig(configForce); err != nil {
				cmd.PrintErrln(err)
			}
			if err := installBinary(); err != nil {
				cmd.PrintErrln(err)
			}
		},
	}

	rootCmd.Flags().BoolP("config-force", "c", false, "If you wish to force write the config")

	return rootCmd
}

// setupHooks initializes and writes the Git hooks.
func setupHooks() error {
	pterm.DefaultSection.Println("Generating .git/hooks")

	pterm.Info.Println("Writing Commit Hooks")
	gitHooks := githooks.NewGitHooks()
	for _, v := range gitHooks.Commit {
		if err := writeConfig(&v); err != nil {
			pterm.Error.Println("✖ Hook write failed for: ", v.Name, err.Error())
			return err
		}
		pterm.Success.Println("✔ Hook written:", v.Name)

	}

	pterm.Info.Println("Writing Checkout Hooks")
	for _, v := range gitHooks.Checkout {
		if err := writeConfig(&v); err != nil {
			pterm.Error.Println("✖ Hook write failed for: ", v.Name, err.Error())
			return err
		}
		pterm.Success.Println("✔ Hook written:", v.Name)
	}

	return nil
}

// writeConfig writes the configuration for a given Git hook using the blueprint package.
func writeConfig(v *githooks.GitHook) error {
	bp, err := blueprint.NewGitHook(v)
	if err != nil {
		return err
	}

	if err := bp.Write(); err != nil {
		return err
	}
	return nil
}

// setupConfig creates and writes the configuration file.
func setupConfig(forceWrite bool) error {
	pterm.DefaultSection.Println("Writing config file")

	config := blueprint.NewBluePrint("railcar.json", "railcar.json", config.RailcarJson, nil)

	if !forceWrite {
		_, err := config.Exists()
		if err == nil {
			pterm.Warning.Println("⚠ Config railcar.json already exists, will not overwrite unless specified")
		} else if errors.Is(err, os.ErrNotExist) {
			forceWrite = true
		}
	}

	if forceWrite {
		if err := config.Write(); err != nil {
			pterm.Error.Println("✖ Failed to write Config railcar.json: ", err.Error())
			return err
		}
		pterm.Success.Println("✔ Config railcar.json successfully written")
	}

	return nil
}

// installBinary writes the embedded binary to the filesystem.
func installBinary() error {
	pterm.DefaultSection.Println("Installing Conductor binary")
	err := embed.WriteBinary()
	if err != nil {
		pterm.Error.Println("✖ Failed to install Conductor: ", err.Error())
		return err
	}
	pterm.Success.Println("✔ Installed conductor successfully")
	return nil
}
