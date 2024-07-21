package assets

import _ "embed"

// This file is meant to embed the parser into skiff. This allows skiff to install the parser onto your machine

//go:embed dist/hookinator_darwin_amd64_v1/hookinator
var MacOSIntel []byte

//go:embed dist/hookinator_darwin_arm64/hookinator
var MacOSSilicon []byte
