package embed

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/devbytes-cloud/hookinator/assets"
)

// Factory pattern

const (
	macOSSilicon string = "darwin-arm64"
	macOSIntel   string = "darwin-amd64"
)

// WriteBinary ...
func WriteBinary() error {
	systemInfo := fmt.Sprintf("%s-%s", fetchOS(), fetchArch())
	binary := fetchBinary(systemInfo)

	if binary == nil {
		panic("couldn't find matching os and arch")
	}

	op := filepath.Join(".", "railcar")
	if err := os.WriteFile(op, binary, 0o755); err != nil {
		return err
	}

	return nil
}

func fetchBinary(systemInfo string) []byte {
	switch systemInfo {
	case macOSSilicon:
		return assets.MacOSSilicon
	case macOSIntel:
		return assets.MacOSIntel
	default:
		return nil
	}
}

// fetchOS returns the current os that is running
func fetchOS() string {
	return runtime.GOOS
}

// fetchArch returns the current architecture that is running
func fetchArch() string {
	return runtime.GOARCH
}
