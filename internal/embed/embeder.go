package embed

import (
	"os"
	"path/filepath"

	"github.com/devbytes-cloud/hookinator/assets"
)

// Factory pattern

// WriteBinary ...
func WriteBinary() error {
	op := filepath.Join(".", "parser")
	if err := os.WriteFile(op, assets.MacOSSilicon, 0o755); err != nil {
		return err
	}

	return nil
}
