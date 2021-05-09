package confy

import (
	"fmt"
	"io"
	"os"
)

var _ Confy = OS("")

// OS is a Confy which uses a filepath to load and store.
type OS string

// Load implements Confy using os.Open and io.ReadAll.
func (o OS) Load() ([]byte, error) {
	f, err := os.Open(string(o))
	if err != nil {
		return nil, fmt.Errorf("open %q: %w", o, err)
	}

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("readall f: %w", err)
	}

	return bs, nil
}

// Store implements Confy using os.Create.
func (o OS) Store(data []byte) error {
	f, err := os.Create(string(o))
	if err != nil {
		return fmt.Errorf("create %q: %w", o, err)
	}

	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("f write: %w", err)
	}

	return nil
}
