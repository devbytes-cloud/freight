package commands

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusCommand(t *testing.T) {
	_, cleanup := withTempGitDir(t)
	defer cleanup()

	t.Run("Status without Freight initialized", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := NewRootCmd()
		cmd.SetOut(&buf)
		cmd.SetArgs([]string{"status"})
		err := cmd.Execute()
		require.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "Freight Status Report")
	})

	t.Run("Status after Freight init", func(t *testing.T) {
		// Initialize Freight
		initCmd := NewRootCmd()
		initCmd.SetArgs([]string{"init", "--allow", "pre-commit"})
		err := initCmd.Execute()
		require.NoError(t, err)

		// Run status
		var buf bytes.Buffer
		statusCmd := NewRootCmd()
		statusCmd.SetOut(&buf)
		statusCmd.SetArgs([]string{"status"})
		err = statusCmd.Execute()
		require.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "pre-commit")
	})

	t.Run("Status with drift", func(t *testing.T) {
		// Initialize Freight
		initCmd := NewRootCmd()
		initCmd.SetArgs([]string{"init", "--allow", "pre-commit"})
		err := initCmd.Execute()
		require.NoError(t, err)

		// Cause drift in pre-commit
		err = os.WriteFile(".git/hooks/pre-commit", []byte("#!/bin/bash\necho drifted"), 0o755)
		require.NoError(t, err)

		// Run status
		var buf bytes.Buffer
		statusCmd := NewRootCmd()
		statusCmd.SetOut(&buf)
		statusCmd.SetArgs([]string{"status"})
		err = statusCmd.Execute()
		require.NoError(t, err)

		output := buf.String()
		assert.Contains(t, output, "drifted")
	})
}
