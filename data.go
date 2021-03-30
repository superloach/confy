package confy

import (
	"io"
)

// Data exposes a storage backend.
type Data interface {
	// Reader fetches a ReadCloser for the given key.
	Reader(key string) (io.ReadCloser, error)

	// Writer fetches a WriteCloser for the given key.
	Writer(key string) (io.WriteCloser, error)

	// Delete removes the given key if it exists.
	Delete(key string) error

	// Keys returns available keys.
	Keys() ([]string, error)
}
