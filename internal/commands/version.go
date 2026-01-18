package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version holds the version number of the application, defaulting to "dev" for development builds.
	Version = "dev"

	// Commit represents the Git commit hash of the build, defaulting to "none" for untracked or local builds.
	Commit = "none"

	// Date represents the build date, defaulting to "unknown" for untracked or local builds.
	Date = "unknown"
)

// versionCommand creates a cobra command for displaying the version of freight.
func versionCommand() *cobra.Command {
	vCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of freight",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(fmt.Sprintf("Freight version: %s", Version))

			verbose, err := cmd.Flags().GetBool("verbose")
			if err != nil {
				cmd.PrintErrf("Failed to get verbose flag: %v", err)
			}

			if verbose {
				cmd.Println(fmt.Sprintf("Commit: %s", Commit))
				cmd.Println(fmt.Sprintf("Date: %s", Date))
			}
		},
	}

	vCmd.Flags().BoolP("verbose", "v", false, "Print verbose version information about freight")

	return vCmd
}
