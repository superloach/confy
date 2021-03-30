package confy

import (
	"fmt"
	"path/filepath"
	"strings"
)

// DefaultExt is the default file extension used for Confy files.
const DefaultExt = ".confy"

// A Confy is a JSON-based configuration directory.
type Confy struct {
	Base string
	Ext  string
}

// Open returns a new Confy with the given base directory.
func Open(base string) (*Confy, error) {
	abase, err := filepath.Abs(base)
	if err != nil {
		return nil, fmt.Errorf("convert %q to absolute path: %w", base, err)
	}

	return &Confy{
		Base: abase,
		Ext:  DefaultExt,
	}, nil
}

func (c *Confy) fp(k string) string {
	return filepath.Join(
		c.Base,
		strings.Join(
			strings.FieldsFunc(
				strings.ToLower(k), // only lowercase
				func(r rune) bool {
					return r < 'a' || r > 'z'
				}, // delimit on non-alpha
			), // split into chunks of alpha
			"_", // rejoin with underscore
		)+c.Ext, // add the extension
	) // join onto base
}
