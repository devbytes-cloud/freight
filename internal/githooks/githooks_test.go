package githooks

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHookPath ensures that hookPath returns the expected path
func TestHookPath(t *testing.T) {
	hookName := "commit-msg"
	expected := filepath.Join(hooksBaseDir, hookName)
	got := hookPath(hookName)
	assert.Equal(t, expected, got, "hookPath should join hooksBaseDir with the hook name")
}

// TestGenerateHooks verifies that generateHooks creates GitHook instances correctly
func TestGenerateHooks(t *testing.T) {
	hookNames := []string{"commit-msg", "post-commit"}
	hooks := generateHooks(hookNames)
	assert.Len(t, hooks, len(hookNames), "generateHooks should return a hook for each input hook name")

	for i, hookName := range hookNames {
		assert.Equal(t, hookName, hooks[i].Name, "hook Name should match the input hook name")
		expectedPath := filepath.Join(hooksBaseDir, hookName)
		assert.Equal(t, expectedPath, hooks[i].Path, "hook Path should be correctly constructed")
		assert.Equal(t, gitHookTemplate, hooks[i].Template, "hook Template should equal gitHookTemplate")
	}
}

// TestGenerateHooksEmpty tests that generateHooks returns an empty slice when given no hook names.
func TestGenerateHooksEmpty(t *testing.T) {
	hooks := generateHooks([]string{})
	assert.Empty(t, hooks, "generateHooks should return an empty slice when no hook names are provided")
}

// TestNewGitHooks verifies that NewGitHooks returns a properly initialized GitHooks structure.
// This test checks that for both commit and checkout hooks the hook paths and templates are configured as expected.
func TestNewGitHooks(t *testing.T) {
	hooksInstance := NewGitHooks()
	assert.NotNil(t, hooksInstance, "NewGitHooks should not return nil")

	// Test commit hooks
	for _, hook := range hooksInstance.Hooks["Commit"] {
		expectedPath := filepath.Join(hooksBaseDir, hook.Name)
		assert.Equal(t, expectedPath, hook.Path, "commit hook Path should match expected")
		assert.Equal(t, gitHookTemplate, hook.Template, "commit hook Template should match gitHookTemplate")
	}

	// Test checkout hooks
	for _, hook := range hooksInstance.Hooks["Checkout"] {
		expectedPath := filepath.Join(hooksBaseDir, hook.Name)
		assert.Equal(t, expectedPath, hook.Path, "checkout hook Path should match expected")
		assert.Equal(t, gitHookTemplate, hook.Template, "checkout hook Template should match gitHookTemplate")
	}
}

// TestAllowedGitHooks verifies that AllowedGitHooks returns the expected map of allowed Git hooks.
func TestAllowedGitHooks(t *testing.T) {
	allowedHooks := AllowedGitHooks()

	expectedHooks := []string{
		PreCommit,
		PrepareCommitMsg,
		CommitMsg,
		PostCommit,
		PostCheckout,
	}

	assert.Len(t, allowedHooks, len(expectedHooks), "AllowedGitHooks should return the correct number of hooks")

	for _, hook := range expectedHooks {
		_, exists := allowedHooks[hook]
		assert.True(t, exists, "Hook %s should be in the allowed hooks map", hook)
	}
}
