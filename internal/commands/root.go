package commands

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/devbytes-cloud/freight/internal/blueprint"
	"github.com/devbytes-cloud/freight/internal/config"
	"github.com/devbytes-cloud/freight/internal/embed"
	"github.com/devbytes-cloud/freight/internal/githooks"
	"github.com/devbytes-cloud/freight/internal/validate"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
		Use:   "freight",
		Short: "Freight is a zero-dependency Git hook manager.",
		Long:  `Freight is a zero-dependency Git hook manager built in Go. It uses a Conductor binary and a Railcar manifest to manage Git hooks across your project.`,
	}

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize Freight in the current repository",
		Long:  `Initialize Freight by installing the Conductor binary, creating a starter Railcar manifest (railcar.json), and rewiring Git hooks.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := validate.GitDirs(); err != nil {
				cmd.PrintErrln(err)
			}

			userAllow, err := cmd.Flags().GetStringSlice("allow")
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			fingerprintPath := ".git/hooks/.fingerprint.yaml"

			var freightConfig config.FreightConfig
			if data, err := os.ReadFile(fingerprintPath); err == nil {
				_ = yaml.Unmarshal(data, &freightConfig)
			}

			var hooksToSetup []string

			if len(userAllow) > 0 {
				hooksToSetup = userAllow
				existingAllow := make(map[string]struct{})
				for _, a := range freightConfig.Allow {
					existingAllow[a] = struct{}{}
				}
				for _, a := range userAllow {
					existingAllow[a] = struct{}{}
				}
				var mergedAllow []string
				for a := range existingAllow {
					mergedAllow = append(mergedAllow, a)
				}
				sort.Strings(mergedAllow)
				freightConfig.Allow = mergedAllow
			} else if len(freightConfig.Allow) == 0 {
				// If no user allow and no existing allow, use default all hooks
				var allHooks []string
				for h := range githooks.AllowedGitHooks() {
					allHooks = append(allHooks, h)
				}
				sort.Strings(allHooks)
				freightConfig.Allow = allHooks
				hooksToSetup = freightConfig.Allow
			} else {
				// If no user allow but existing allow exists, use existing allow
				hooksToSetup = freightConfig.Allow
			}

			pterm.Debug.Printfln("Effective allow: %v", freightConfig.Allow)
			pterm.Debug.Printfln("Hooks to setup: %v", hooksToSetup)

			validatedAllow, err := validate.GitHooks(hooksToSetup)
			if err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
			}

			if err := setupHooks(validatedAllow); err != nil {
				cmd.PrintErrln(err)
				os.Exit(1)
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

			// Save fingerprint
			freightConfig.Version = Version
			data, err := yaml.Marshal(freightConfig)
			if err == nil {
				pterm.DefaultSection.Println("Writing fingerprint file")
				comment := "# This file is managed by Freight. It keeps track of the version and allowed hooks.\n"
				finalData := append([]byte(comment), data...)
				if err := os.WriteFile(fingerprintPath, finalData, 0o644); err != nil {
					pterm.Error.Printfln("✖ Failed to write fingerprint: %v", err)
				} else {
					pterm.Success.Printfln("✔ Fingerprint .fingerprint.yaml successfully written to .git/hooks")
				}
			}

			pterm.Success.Println("Freight initialized successfully!")
		},
	}

	initCmd.Flags().BoolP("config-force", "c", false, "If you wish to force write the config")
	initCmd.Flags().StringSliceP("allow", "a", []string{}, "Specific Git hooks to install (default: all). Valid options: pre-commit, prepare-commit-msg, commit-msg, post-commit, post-checkout")
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(statusCommand())
	rootCmd.AddCommand(versionCommand())

	return rootCmd
}

// setupHooks initializes and writes the Git hooks to the .git/hooks directory.
// It only writes the hooks that are included in the allowedHooks map.
func setupHooks(allowedHooks map[string]struct{}) error {
	pterm.DefaultSection.Println("Generating .git/hooks")
	pterm.Debug.Printfln("Allowed hooks: %v", allowedHooks)

	gitHooks := githooks.NewGitHooks()
	for hookName, hookGroup := range gitHooks.Hooks {
		pterm.Info.Println("Writing", hookName, "Hooks")
		for _, v := range hookGroup {
			if _, ok := allowedHooks[v.Name]; ok {
				if err := writeConfig(&v); err != nil {
					pterm.Error.Println("✖ Hook write failed for: ", v.Name, err.Error())
					return err
				}
				pterm.Success.Println("✔ Hook written:", v.Name)
			} else {
				pterm.Warning.Println("Skipping hook:", v.Name, "not allowed")
			}
		}
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

// setupConfig initializes the Railcar configuration file (railcar.json).
// If forceWrite is true, it overwrites any existing configuration.
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

// installBinary extracts and writes the embedded Conductor binary to the current directory.
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
