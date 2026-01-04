package assets

import _ "embed"

// This file is meant to embed the conductor into freight. This allows freight to install the conductor onto your machine

//go:embed dist/conductor_darwin_amd64_v1/conductor
var MacOSIntel []byte

//go:embed dist/conductor_darwin_arm64_v8.0/conductor
var MacOSSilicon []byte

//go:embed dist/conductor_linux_amd64_v1/conductor
var LinuxAMD64 []byte

//go:embed dist/conductor_linux_arm64_v8.0/conductor
var LinuxARM64 []byte

//go:embed dist/conductor_linux_arm_6/conductor
var LinuxARM32 []byte

//go:embed dist/conductor_windows_amd64_v1/conductor.exe
var WindowsAMD64 []byte

//go:embed dist/conductor_windows_arm64_v8.0/conductor.exe
var WindowsARM64 []byte
