package confy

import (
	"io"
)

// Form describes an object encoding format.
type Form interface {
	// Marshal encodes the given value into the given Writer.
	Marshal(w io.Writer, value interface{}) error

	// Unmarshal decodes the given Reader into the given pointer.
	Unmarshal(r io.Reader, ptr interface{}) error
}
