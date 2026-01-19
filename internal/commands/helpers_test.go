package commands

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func withTempGitDir(t *testing.T) (string, func()) {
	t.Helper()

	tmpDir, err := os.MkdirTemp("", "freight-test-*")
	require.NoError(t, err)

	origWd, err := os.Getwd()
	require.NoError(t, err)

	err = os.Chdir(tmpDir)
	require.NoError(t, err)

	err = os.MkdirAll(".git/hooks", 0o755)
	require.NoError(t, err)

	cleanup := func() {
		os.Chdir(origWd)
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}
