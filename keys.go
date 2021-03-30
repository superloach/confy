package confy

import (
	"fmt"
	"os"
	"strings"
)

// Keys returns available keys in the Confy, given by regular files matching its Ext.
func (c *Confy) Keys() ([]string, error) {
	ds, err := os.ReadDir(c.Base)
	if err != nil {
		return nil, fmt.Errorf("read dir %q: %w", c.Base, err)
	}

	ks := make([]string, 0, len(ds))

	for _, d := range ds {
		if !d.Type().IsRegular() {
			continue
		}

		n := d.Name()
		if strings.HasSuffix(n, c.Ext) {
			ks = append(ks, strings.TrimSuffix(n, c.Ext))
		}
	}

	return ks, nil
}
