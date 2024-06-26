package hookinator

import _ "embed"

// This file is meant to embed the parser into skiff. This allows skiff to install the parser onto your machine

//go:embed assets/parser
var Parser []byte
