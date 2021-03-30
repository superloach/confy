package confy

import (
	"fmt"
	"os"
)

// Del removes the given key and its value from the Confy.
func (c *Confy) Del(key string) error {
	fp := c.fp(key)

	if err := os.Remove(fp); err != nil {
		return fmt.Errorf("remove %q: %w", fp, err)
	}

	return nil
}
