package validate_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/devbytes-cloud/freight/internal/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitDirs(t *testing.T) {
	tt := []struct {
		name  string
		fail  bool
		paths []string
	}{
		{
			name:  "success",
			fail:  false,
			paths: []string{".git", "hooks"},
		},
		{
			name:  "missing .git",
			fail:  true,
			paths: []string{},
		},
		{
			name:  "missing .git hooks",
			fail:  true,
			paths: []string{".git"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			require.NoError(t, os.MkdirAll(filepath.Join(append([]string{tempDir}, tc.paths...)...), 0o755))

			oldWd, err := os.Getwd()
			require.NoError(t, err)
			require.NoError(t, os.Chdir(tempDir))
			defer func() {
				_ = os.Chdir(oldWd)
			}()

			err = validate.GitDirs()
			if tc.fail {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCurrentWD(t *testing.T) {
	wd, err := validate.CurrentWD()
	assert.NoError(t, err)
	assert.NotEmpty(t, wd)
}
