package commands

import (
	_ "embed"
	"errors"
	"fmt"
	"os"

	"github.com/devbytes-cloud/freight/internal/blueprint"
	"github.com/devbytes-cloud/freight/internal/embed"
	"github.com/devbytes-cloud/freight/internal/githooks"
	"github.com/devbytes-cloud/freight/internal/validate"
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
			if err := validate.GitDirExists(); err != nil {
				panic(err)
			}
			if err := setupHooks(); err != nil {
				panic(err)
			}

			configForce, err := cmd.Flags().GetBool("config-force")
			if err != nil {
				panic(err)
			}

			if err := setupConfig(configForce); err != nil {
				panic(err)
			}
			if err := installBinary(); err != nil {
				panic(err)
			}
		},
	}

	rootCmd.Flags().BoolP("config-force", "c", false, "If you wish to force write the config")

	return rootCmd
}

// setupHooks initializes and writes the Git hooks.
func setupHooks() error {
	fmt.Println("Generating githook files")
	fmt.Println("=========================")

	gitHooks := githooks.NewGitHooks()
	for _, v := range gitHooks.Commit {
		bp, err := blueprint.NewGitHook(&v)
		if err != nil {
			return err
		}

		if err := bp.Write(); err != nil {
			return err
		}
	}

	return nil
}

// setupConfig creates and writes the configuration file.
func setupConfig(forceWrite bool) error {
	fmt.Println("Generating config file")
	fmt.Println("=========================")

	config := blueprint.NewConfig()

	if !forceWrite {
		if _, err := config.Exists(); err != nil {
			if errors.Is(err, os.ErrNotExist) {
				forceWrite = true
			}
		}
	}

	if forceWrite {
		if err := blueprint.NewConfig().Write(); err != nil {
			return err
		}
	}

	return nil
}

// installBinary writes the embedded binary to the filesystem.
func installBinary() error {
	fmt.Println("Generating freight binary")
	fmt.Println("=========================")
	if err := embed.WriteBinary(); err != nil {
		return err
	}
	return nil
}
