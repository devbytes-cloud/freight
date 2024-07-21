package assets

import _ "embed"

// This file is meant to embed the railcar into skiff. This allows skiff to install the railcar onto your machine

//go:embed dist/railcar_darwin_amd64_v1/railcar
var MacOSIntel []byte

//go:embed dist/railcar_darwin_arm64/railcar
var MacOSSilicon []byte
