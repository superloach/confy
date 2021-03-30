package confy

import (
	"fmt"
	"io"
)

// MustForm can be used to wrap functions that return (Form, error), panicking on error.
func MustForm(f Form, err error) Form {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}

	return f
}

// Form describes an object encoding format.
type Form interface {
	// Marshal encodes the given value into the given Writer.
	Marshal(w io.Writer, value interface{}) error

	// Unmarshal decodes the given Reader into the given pointer.
	Unmarshal(r io.Reader, ptr interface{}) error
}
