package assets

import _ "embed"

// This file is meant to embed the railcar into freight. This allows freight to install the railcar onto your machine

//go:embed dist/railcar_darwin_amd64_v1/railcar
var MacOSIntel []byte

//go:embed dist/railcar_darwin_arm64/railcar
var MacOSSilicon []byte

//go:embed dist/railcar_linux_amd64_v1/railcar
var LinuxAMD64 []byte

//go:embed dist/railcar_linux_arm64/railcar
var LinuxARM64 []byte

//go:embed dist/railcar_linux_arm_6/railcar
var LinuxARM32 []byte

//go:embed dist/railcar_windows_amd64_v1/railcar.exe
var WindowsAMD64 []byte

//go:embed dist/railcar_windows_arm64/railcar.exe
var WindowsARM64 []byte

//go:embed dist/railcar_windows_arm_6/railcar.exe
var WindowsARM32 []byte
