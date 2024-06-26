package embed

import (
	"os"
	"path/filepath"

	"github.com/devbytes-cloud/hookinator"
)

// WriteBinary ...
func WriteBinary() error {
	op := filepath.Join(".", "parser")
	if err := os.WriteFile(op, hookinator.Parser, 0o755); err != nil {
		return err
	}

	return nil
}
