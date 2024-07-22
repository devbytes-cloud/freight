package embed

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteBinary(t *testing.T) {
	err := WriteBinary()
	assert.NoError(t, err)
	assert.FileExists(t, "./railcar")
	defer func() {
		err := os.Remove("./railcar")
		assert.NoError(t, err)
	}()
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
