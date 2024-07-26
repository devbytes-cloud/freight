package embed

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/devbytes-cloud/freight/assets"
)

// List of supported OS + Arch railcar binaries
const (
	macOSSilicon string = "darwin-arm64"
	macOSIntel   string = "darwin-amd64"
	linuxAMD64   string = "linux-amd64"
	linuxARM64   string = "linux-arm64"
	linuxARM32   string = "linux-arm"
	windowsAMD64 string = "windows-amd64"
	windowsARM64 string = "windows-arm64"
	windowsARM32 string = "windows-arm"
)

// WriteBinary will install railcar into your working directory
func WriteBinary() error {
	systemInfo := fmt.Sprintf("%s-%s", fetchOS(), fetchArch())
	binary := fetchBinary(systemInfo)

	if binary == nil {
		return fmt.Errorf("no matching railcar binary for %s", systemInfo)
	}

	op := filepath.Join(".", "railcar")
	return os.WriteFile(op, binary, 0o755)
}

// fetchBinary will return the proper railcar binary for your system
func fetchBinary(systemInfo string) []byte {
	switch systemInfo {
	case macOSSilicon:
		return assets.MacOSSilicon
	case macOSIntel:
		return assets.MacOSIntel
	case linuxAMD64:
		return assets.LinuxAMD64
	case linuxARM64:
		return assets.LinuxARM64
	case linuxARM32:
		return assets.LinuxARM32
	case windowsAMD64:
		return assets.WindowsAMD64
	case windowsARM64:
		return assets.WindowsARM64
	case windowsARM32:
		return assets.WindowsARM32
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
