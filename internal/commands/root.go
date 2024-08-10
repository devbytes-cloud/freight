package commands

import (
	_ "embed"
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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// rootCmd is the entry command into freight
var rootCmd = &cobra.Command{
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
		if err := setupConfig(); err != nil {
			panic(err)
		}
		if err := installBinary(); err != nil {
			panic(err)
		}
	},
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
func setupConfig() error {
	fmt.Println("Generating config file")
	fmt.Println("=========================")
	if err := blueprint.NewConfig().Write(); err != nil {
		return err
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
