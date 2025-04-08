package embed

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteBinary(t *testing.T) {
	tmpDir := t.TempDir()

	oldWD, err := os.Getwd()
	require.NoError(t, err)
	defer func(dir string) {
		require.NoError(t, os.Chdir(dir))
	}(oldWD)

	require.NoError(t, os.Chdir(tmpDir))

	assert.NoError(t, WriteBinary())
	assert.FileExists(t, "./conductor")
}

func TestFetchBinary(t *testing.T) {
	t.Run("binary does not exist", func(t *testing.T) {
		assert.Nil(t, fetchBinary("junk"))
	})

	t.Run("binary exists", func(t *testing.T) {
		systemInfo := fmt.Sprintf("%s-%s", fetchOS(), fetchArch())
		assert.NotNil(t, fetchBinary(systemInfo))
	})
}

func TestFetchOS(t *testing.T) {
	assert.Equal(t, runtime.GOOS, fetchOS())
}

func TestFetchArch(t *testing.T) {
	assert.Equal(t, runtime.GOARCH, fetchArch())
}
