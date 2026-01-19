package commands

import (
	"os"
	"sort"
	"testing"

	"github.com/devbytes-cloud/freight/internal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestInitMergeAllow(t *testing.T) {
	_, cleanup := withTempGitDir(t)
	defer cleanup()

	fingerprintPath := ".git/hooks/.fingerprint.yaml"

	t.Run("Initial init with specific allow", func(t *testing.T) {
		cmd := NewRootCmd()
		cmd.SetArgs([]string{"init", "--allow", "pre-commit,commit-msg"})
		err := cmd.Execute()
		require.NoError(t, err)

		// Check fingerprint
		data, err := os.ReadFile(fingerprintPath)
		require.NoError(t, err)

		var cfg config.FreightConfig
		err = yaml.Unmarshal(data, &cfg)
		require.NoError(t, err)

		sort.Strings(cfg.Allow)
		assert.Equal(t, []string{"commit-msg", "pre-commit"}, cfg.Allow)
	})

	t.Run("Merge new allow with existing", func(t *testing.T) {
		cmd := NewRootCmd()
		cmd.SetArgs([]string{"init", "--allow", "post-commit"})
		err := cmd.Execute()
		require.NoError(t, err)

		// Check fingerprint
		data, err := os.ReadFile(fingerprintPath)
		require.NoError(t, err)

		var cfg config.FreightConfig
		err = yaml.Unmarshal(data, &cfg)
		require.NoError(t, err)

		sort.Strings(cfg.Allow)
		assert.Equal(t, []string{"commit-msg", "post-commit", "pre-commit"}, cfg.Allow)
	})

	t.Run("No duplicates when merging", func(t *testing.T) {
		cmd := NewRootCmd()
		cmd.SetArgs([]string{"init", "--allow", "pre-commit,post-checkout"})
		err := cmd.Execute()
		require.NoError(t, err)

		// Check fingerprint
		data, err := os.ReadFile(fingerprintPath)
		require.NoError(t, err)

		var cfg config.FreightConfig
		err = yaml.Unmarshal(data, &cfg)
		require.NoError(t, err)

		sort.Strings(cfg.Allow)
		assert.Equal(t, []string{"commit-msg", "post-checkout", "post-commit", "pre-commit"}, cfg.Allow)
	})

	t.Run("Only specified hooks are initialized with --allow", func(t *testing.T) {
		// Clean up and start fresh
		os.RemoveAll(".git/hooks")
		os.MkdirAll(".git/hooks", 0o755)

		// 1. Init with pre-commit
		cmd := NewRootCmd()
		cmd.SetArgs([]string{"init", "--allow", "pre-commit"})
		err := cmd.Execute()
		require.NoError(t, err)

		require.FileExists(t, ".git/hooks/pre-commit")
		_, err = os.Stat(".git/hooks/post-commit")
		assert.True(t, os.IsNotExist(err))

		// 2. Init with post-commit. It should NOT re-initialize pre-commit (though pre-commit should still exist if it was there)
		// To truly test it ONLY runs post-commit, we can delete pre-commit and see if it comes back.
		os.Remove(".git/hooks/pre-commit")

		cmd = NewRootCmd()
		cmd.SetArgs([]string{"init", "--allow", "post-commit"})
		err = cmd.Execute()
		require.NoError(t, err)

		require.FileExists(t, ".git/hooks/post-commit")
		_, err = os.Stat(".git/hooks/pre-commit")
		assert.True(t, os.IsNotExist(err), "pre-commit should not have been re-initialized")

		// 3. Check fingerprint has BOTH
		data, err := os.ReadFile(fingerprintPath)
		require.NoError(t, err)
		var cfg config.FreightConfig
		err = yaml.Unmarshal(data, &cfg)
		require.NoError(t, err)
		sort.Strings(cfg.Allow)
		assert.Equal(t, []string{"post-commit", "pre-commit"}, cfg.Allow)
	})
}
