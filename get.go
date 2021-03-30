package confy

import (
	"encoding/json"
	"fmt"
	"os"
)

// Get loads a value for the given key from the Confy, and stores it into the given value pointer.
func (c *Confy) Get(key string, value interface{}) error {
	fp := c.fp(key)

	f, err := os.Open(fp)
	if err != nil {
		return fmt.Errorf("open %q: %w", fp, err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(value)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	return nil
}
