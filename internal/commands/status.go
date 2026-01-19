package commands

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"sort"

	"github.com/devbytes-cloud/freight/internal/config"
	"github.com/devbytes-cloud/freight/internal/githooks"
	"github.com/devbytes-cloud/freight/internal/validate"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// statusCommand returns a cobra.Command that reports the current state of Freight in the repository.
func statusCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Report the current state of Freight in the repository",
		Run: func(cmd *cobra.Command, args []string) {
			pterm.SetDefaultOutput(cmd.OutOrStdout())
			pterm.DefaultHeader.Println("Freight Status Report")

			// Check Git repository
			gitErr := validate.GitDirs()
			if gitErr == nil {
				pterm.Success.Println("Valid Git repository")
			} else {
				pterm.Error.Println("Not a valid Git repository")
			}

			railcarPath := "railcar.json"
			railcarExists := fileExists(railcarPath)
			printStatus("Railcar Manifest (railcar.json)", railcarExists)

			fingerprintPath := ".git/hooks/.fingerprint.yaml"
			fingerprintExists := fileExists(fingerprintPath)
			printStatus("Fingerprint (.git/hooks/.fingerprint.yaml)", fingerprintExists)

			conductorExists := fileExists("conductor")
			printStatus("Conductor Binary", conductorExists)

			var freightConfig config.FreightConfig
			if fingerprintExists {
				data, err := os.ReadFile(fingerprintPath)
				if err == nil {
					_ = yaml.Unmarshal(data, &freightConfig)
				}
			}

			if freightConfig.Version != "" {
				pterm.Info.Printfln("Last Applied Freight Version: %s", freightConfig.Version)
			} else {
				pterm.Info.Println("Last Applied Freight Version: Unknown")
			}

			pterm.Info.Printfln("Current Freight Version: %s", Version)

			pterm.DefaultSection.Println("Git Hooks Status")

			hooksData := pterm.TableData{
				{"Hook", "Exists", "Managed by Freight", "Drift"},
			}

			managedHooks := make(map[string]bool)
			for _, hook := range freightConfig.Allow {
				managedHooks[hook] = true
			}

			gitHooks := githooks.NewGitHooks()
			hookMap := make(map[string]githooks.GitHook)
			var allHooks []string
			for _, hookGroup := range gitHooks.Hooks {
				for _, h := range hookGroup {
					hookMap[h.Name] = h
					allHooks = append(allHooks, h.Name)
				}
			}
			sort.Strings(allHooks)

			wd, _ := validate.CurrentWD()

			for _, hookName := range allHooks {
				hookPath := filepath.Join(".git/hooks", hookName)
				exists := fileExists(hookPath)
				managed := managedHooks[hookName]
				drift := "N/A"

				if managed {
					if !exists {
						drift = "missing"
					} else {
						// Drift detection
						gh, ok := hookMap[hookName]
						if ok {
							expected, err := renderTemplate(gh, wd)
							if err != nil {
								drift = "error"
							} else {
								actual, err := os.ReadFile(hookPath)
								if err != nil {
									drift = "error"
								} else if bytes.Equal(actual, []byte(expected)) {
									drift = "none"
								} else {
									drift = "drifted"
								}
							}
						}
					}
				}

				hooksData = append(hooksData, []string{
					hookName,
					boolToCheck(exists),
					boolToCheck(managed),
					drift,
				})
			}

			pterm.DefaultTable.WithHasHeader().WithData(hooksData).Render()
		},
	}

	return cmd
}

// fileExists checks if a file exists at the given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// printStatus prints a status message for a given component, indicating whether it was found or is missing.
func printStatus(name string, exists bool) {
	if exists {
		pterm.Info.Printfln("%s: Found", name)
	} else {
		pterm.Error.Printfln("%s: Missing", name)
	}
}

// boolToCheck converts a boolean value to a checkmark (✔) or a cross (✘) symbol.
func boolToCheck(b bool) string {
	if b {
		return "✔"
	}
	return "✘"
}

// renderTemplate renders the Git hook template with the provided workspace directory and hook information.
func renderTemplate(gh githooks.GitHook, wd string) (string, error) {
	tmpl, err := template.New(gh.Name).Parse(gh.Template)
	if err != nil {
		return "", err
	}

	path := struct {
		Path string
		Type string
	}{
		Path: wd,
		Type: gh.Name,
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, path); err != nil {
		return "", err
	}

	return buf.String(), nil
}
