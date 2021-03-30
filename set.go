package confy

import (
	"encoding/json"
	"fmt"
	"os"
)

// Set stores the given value into the given key in the Confy.
func (c *Confy) Set(key string, value interface{}) error {
	fp := c.fp(key)

	f, err := os.Create(fp)
	if err != nil {
		return fmt.Errorf("create %q: %w", fp, err)
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(value)
	if err != nil {
		return fmt.Errorf("json encode %v: %w", value, err)
	}

	return nil
}
