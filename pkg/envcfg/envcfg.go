package envcfg

import (
	"context"
	"os"
	"strings"

	"github.com/heetch/confita/backend"
)

// NewBackend creates a configuration loader that loads from the environment.
func NewBackend(prefix string) backend.Backend {
	return backend.Func("env", func(ctx context.Context, key string) ([]byte, error) {
		key = prefix + "_" + strings.Replace(strings.ToUpper(key), "-", "_", -1)
		if val := os.Getenv(key); val != "" {
			return []byte(val), nil
		}
		return nil, backend.ErrNotFound
	})
}
