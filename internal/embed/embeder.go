package embed

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/devbytes-cloud/freight/assets"
)

// List of supported OS + Arch conductor binaries
const (
	// macOSSilicon represents the macOS Silicon architecture (darwin-arm64).
	macOSSilicon string = "darwin-arm64"
	// macOSIntel represents the macOS Intel architecture (darwin-amd64).
	macOSIntel string = "darwin-amd64"
	// linuxAMD64 represents the Linux AMD64 architecture (linux-amd64).
	linuxAMD64 string = "linux-amd64"
	// linuxARM64 represents the Linux ARM64 architecture (linux-arm64).
	linuxARM64 string = "linux-arm64"
	// linuxARM32 represents the Linux ARM32 architecture (linux-arm).
	linuxARM32 string = "linux-arm"
	// windowsAMD64 represents the Windows AMD64 architecture (windows-amd64).
	windowsAMD64 string = "windows-amd64"
	// windowsARM64 represents the Windows ARM64 architecture (windows-arm64).
	windowsARM64 string = "windows-arm64"
)

// WriteBinary will install conductor into your working directory
func WriteBinary() error {
	systemInfo := fmt.Sprintf("%s-%s", fetchOS(), fetchArch())
	binary := fetchBinary(systemInfo)

	if binary == nil {
		return fmt.Errorf("no matching conductor binary for %s", systemInfo)
	}

	op := filepath.Join(".", "conductor")

	if err := os.WriteFile(op, binary, 0o755); err != nil {
		return fmt.Errorf("error writing conductor binary: %s", err)
	}
	return nil
}

// fetchBinary will return the proper conductor binary for your system
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
