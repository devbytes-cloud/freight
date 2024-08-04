package commands

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/devbytes-cloud/freight/internal/githooks"

	"github.com/devbytes-cloud/freight/internal/validate"

	"github.com/devbytes-cloud/freight/internal/blueprint"
	"github.com/devbytes-cloud/freight/internal/embed"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "init",
	Long:  `init`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cheaing a bit but i am doing init here")

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

func setupHooks() error {
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

func setupConfig() error {
	if err := blueprint.NewConfig().Write(); err != nil {
		return err
	}
	return nil
}

func installBinary() error {
	if err := embed.WriteBinary(); err != nil {
		return err
	}
	return nil
}
